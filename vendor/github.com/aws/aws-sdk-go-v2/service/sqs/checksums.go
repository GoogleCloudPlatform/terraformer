package sqs

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	request "github.com/aws/aws-sdk-go-v2/aws"
)

var (
	errChecksumMissingBody = fmt.Errorf("cannot compute checksum. missing body")
	errChecksumMissingMD5  = fmt.Errorf("cannot verify checksum. missing response MD5")
)

func setupChecksumValidation(r *request.Request) {
	switch r.Operation.Name {
	case opSendMessage:
		r.Handlers.Unmarshal.PushBack(verifySendMessage)
	case opSendMessageBatch:
		r.Handlers.Unmarshal.PushBack(verifySendMessageBatch)
	case opReceiveMessage:
		r.Handlers.Unmarshal.PushBack(verifyReceiveMessage)
	}
}

func verifySendMessage(r *request.Request) {
	if r.ParamsFilled() {
		in := r.Params.(*SendMessageInput)
		out := r.Data.(*SendMessageOutput)
		err := checksumsMatch(in.MessageBody, out.MD5OfMessageBody)
		if err != nil {
			r.Error = &InvalidChecksumError{
				Reason: "response checksum does not match sent",
				Err:    err,
			}
		}
	}
}

func verifySendMessageBatch(r *request.Request) {
	if r.ParamsFilled() {
		entries := map[string]SendMessageBatchResultEntry{}
		ids := []string{}

		out := r.Data.(*SendMessageBatchOutput)
		for _, entry := range out.Successful {
			entries[*entry.Id] = entry
		}

		in := r.Params.(*SendMessageBatchInput)
		for _, entry := range in.Entries {
			if e, ok := entries[*entry.Id]; ok {
				if err := checksumsMatch(entry.MessageBody, e.MD5OfMessageBody); err != nil {
					ids = append(ids, *e.MessageId)
				}
			}
		}
		if len(ids) > 0 {
			r.Error = &InvalidChecksumError{
				Reason: fmt.Sprintf("invalid messages: %s", strings.Join(ids, ", ")),
			}
		}
	}
}

func verifyReceiveMessage(r *request.Request) {
	if r.ParamsFilled() {
		ids := []string{}
		out := r.Data.(*ReceiveMessageOutput)
		for i, msg := range out.Messages {
			err := checksumsMatch(msg.Body, msg.MD5OfBody)
			if err != nil {
				if msg.MessageId == nil {
					if r.Config.Logger != nil {
						r.Config.Logger.Log(fmt.Sprintf(
							"WARN: SQS.ReceiveMessage failed checksum request id: %s, message %d has no message ID.",
							r.RequestID, i,
						))
					}
					continue
				}

				ids = append(ids, *msg.MessageId)
			}
		}
		if len(ids) > 0 {
			r.Error = &InvalidChecksumError{
				Reason: fmt.Sprintf("invalid messages: %s", strings.Join(ids, ", ")),
			}
		}
	}
}

func checksumsMatch(body, expectedMD5 *string) error {
	if body == nil {
		return errChecksumMissingBody
	} else if expectedMD5 == nil {
		return errChecksumMissingMD5
	}

	msum := md5.Sum([]byte(*body))
	sum := hex.EncodeToString(msum[:])
	if sum != *expectedMD5 {
		return fmt.Errorf("expected MD5 checksum '%s', got '%s'", *expectedMD5, sum)
	}

	return nil
}

// InvalidChecksumError provides the error type for invalid checksum errors.
type InvalidChecksumError struct {
	Reason string
	Err    error
}

// RetryableError decorates this error as retryable.
func (e *InvalidChecksumError) RetryableError() bool { return true }

// Unwrap returns the underlying error if there was one.
func (e *InvalidChecksumError) Unwrap() error { return e.Err }

func (e *InvalidChecksumError) Error() string {
	var extra string
	if e.Err != nil {
		extra = ", " + e.Err.Error()
	}
	return fmt.Sprintf("checksum validation error, %v%s", e.Reason, extra)
}
