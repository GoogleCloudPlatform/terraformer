package kubernetes

import (
	"fmt"
	"log"
	"sort"

	api "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	kubernetes "k8s.io/client-go/kubernetes"
)

func getLastWarningsForObject(conn *kubernetes.Clientset, metadata meta_v1.ObjectMeta, kind string, limit int) ([]api.Event, error) {
	m := map[string]string{
		"involvedObject.name": metadata.Name,
		"involvedObject.kind": kind,
	}
	if metadata.Namespace != "" {
		m["involvedObject.namespace"] = metadata.Namespace
	}

	fs := fields.Set(m).String()
	log.Printf("[DEBUG] Looking up events via this selector: %q", fs)
	out, err := conn.CoreV1().Events(metadata.Namespace).List(meta_v1.ListOptions{
		FieldSelector: fs,
	})
	if err != nil {
		return nil, err
	}

	// It would be better to sort & filter on the server-side
	// but API doesn't seem to support it
	var warnings []api.Event

	// Bring latest events to the top, for easy access
	sort.Slice(out.Items, func(i, j int) bool {
		return out.Items[i].LastTimestamp.After(out.Items[j].LastTimestamp.Time)
	})

	log.Printf("[DEBUG] Received %d events for %s/%s (%s)",
		len(out.Items), metadata.Namespace, metadata.Name, kind)

	warnCount := 0
	uniqueWarnings := make(map[string]api.Event, 0)
	for _, e := range out.Items {
		if warnCount >= limit {
			break
		}

		if e.Type == api.EventTypeWarning {
			_, found := uniqueWarnings[e.Message]
			if found {
				continue
			}
			warnings = append(warnings, e)
			uniqueWarnings[e.Message] = e
			warnCount++
		}
	}

	return warnings, nil
}

func stringifyEvents(events []api.Event) string {
	var output string
	for _, e := range events {
		output += fmt.Sprintf("\n   * %s (%s): %s: %s",
			e.InvolvedObject.Name, e.InvolvedObject.Kind,
			e.Reason, e.Message)
	}
	return output
}
