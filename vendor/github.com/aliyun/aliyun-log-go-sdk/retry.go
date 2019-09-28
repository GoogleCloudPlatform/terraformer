package sls

import (
	"golang.org/x/net/context"

	"github.com/cenkalti/backoff"
	"github.com/pkg/errors"
)

// Retry execute the input operation immediately at first,
// and do an exponential backoff retry when failed.
// The default max elapsed time is 15 minutes.
// The default retry intervals are shown below, in seconds.
//  1          0.5                     [0.25,   0.75]
//  2          0.75                    [0.375,  1.125]
//  3          1.125                   [0.562,  1.687]
//  4          1.687                   [0.8435, 2.53]
//  5          2.53                    [1.265,  3.795]
//  6          3.795                   [1.897,  5.692]
//  7          5.692                   [2.846,  8.538]
//  8          8.538                   [4.269, 12.807]
//  9         12.807                   [6.403, 19.210]
// ...
// The signature of backoff.Operation is "func() error".
func Retry(ctx context.Context, o backoff.Operation) error {
	return RetryWithBackOff(ctx, backoff.NewExponentialBackOff(), o)
}

// RetryWithBackOff ...
func RetryWithBackOff(ctx context.Context, b backoff.BackOff, o backoff.Operation) error {
	ticker := backoff.NewTicker(b)
	defer ticker.Stop()
	var err error
	for {
		select {
		case <-ctx.Done():
			return errors.Wrapf(ctx.Err(), "stopped retrying err: %v", err)
		default:
			select {
			case _, ok := <-ticker.C:
				if !ok {
					return err
				}
				err = o()
				if err == nil {
					return nil
				}
			case <-ctx.Done():
				return errors.Wrapf(ctx.Err(), "stopped retrying err: %v", err)
			}
		}
	}
}

// ConditionOperation : retry depends on the retured bool
type ConditionOperation func() (bool, error)

// RetryWithCondition ...
func RetryWithCondition(ctx context.Context, b backoff.BackOff, o ConditionOperation) error {
	ticker := backoff.NewTicker(b)
	defer ticker.Stop()
	var err error
	var needRetry bool
	for {
		select {
		case <-ctx.Done():
			return errors.Wrapf(ctx.Err(), "stopped retrying err: %v", err)
		default:
			select {
			case _, ok := <-ticker.C:
				if !ok {
					return err
				}
				needRetry, err = o()
				if !needRetry {
					return err
				}
			case <-ctx.Done():
				return errors.Wrapf(ctx.Err(), "stopped retrying err: %v", err)
			}
		}
	}
}

// RetryWithAttempt ...
func RetryWithAttempt(ctx context.Context, maxAttempt int, o ConditionOperation) error {
	b := backoff.NewExponentialBackOff()
	ticker := backoff.NewTicker(b)
	defer ticker.Stop()
	var err error
	var needRetry bool
	for try := 0; try < maxAttempt; try++ {
		// make sure we check ctx.Done() first
		select {
		case <-ctx.Done():
			return errors.Wrapf(ctx.Err(), "stopped retrying err: %v", err)
		default:
		}

		select {
		case _, ok := <-ticker.C:
			if !ok {
				return err
			}
			needRetry, err = o()
			if !needRetry {
				return err
			}
		}
	}
	return err
}
