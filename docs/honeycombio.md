### Use with Honeycomb.io

Example:

```
export HONEYCOMB_API_KEY=MYAPIKEY
./terraformer import honeycombio --resources=board,trigger
```

List of supported Honeycomb resources:

* `board`
  * `honeycombio_board`
  * `honeycombio_query`
  * `honeycombio_query_annotation`
* `derived_column`
  * `honeycombio_derived_column`
* `trigger`
  * `honeycombio_trigger`
  * `honeycombio_query`
* `slo`
  * `honeycombio_slo`
  * `honeycombio_burn_alert`
  * `honeycombio_derived_column`
