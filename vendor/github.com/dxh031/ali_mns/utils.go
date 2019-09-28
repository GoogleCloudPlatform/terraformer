package ali_mns

import (
	"bytes"

	"github.com/gogap/errors"
	"github.com/valyala/fasthttp"
)

func send(client MNSClient, decoder MNSDecoder, method Method, headers map[string]string, message interface{}, resource string, v interface{}) (statusCode int, err error) {
	var resp *fasthttp.Response
	if resp, err = client.Send(method, headers, message, resource); err != nil {
		return
	}

	if resp != nil {
		statusCode = resp.Header.StatusCode()

		if statusCode != fasthttp.StatusCreated &&
			statusCode != fasthttp.StatusOK &&
			statusCode != fasthttp.StatusNoContent {

			// get the response body
			//   the body is set in error when decoding xml failed
			bodyBytes := resp.Body()

			var e2 error
			err, e2 = decoder.DecodeError(bodyBytes, resource)

			if e2 != nil {
				err = ERR_UNMARSHAL_ERROR_RESPONSE_FAILED.New(errors.Params{"err": e2, "resp": string(bodyBytes)})
				return
			}
			return
		}

		if v != nil {
			buf := bytes.NewReader(resp.Body())
			if e := decoder.Decode(buf, v); e != nil {
				err = ERR_UNMARSHAL_RESPONSE_FAILED.New(errors.Params{"err": e})
				return
			}
		}
	}

	return
}
