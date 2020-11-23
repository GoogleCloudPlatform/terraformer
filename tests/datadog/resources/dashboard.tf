resource "datadog_dashboard" "ordered_dashboard_example" {
  title         = "Ordered Layout Dashboard"
  description   = "Created using the Datadog provider in Terraform"
  layout_type   = "ordered"
  is_read_only  = true

  widget {
    alert_graph_definition {
      alert_id = "895605"
      viz_type = "timeseries"
      title = "Widget Title"
      time = {
        live_span = "1h"
      }
    }
  }

  widget {
    alert_value_definition {
      alert_id = "895605"
      precision = 3
      unit = "b"
      text_align = "center"
      title = "Widget Title"
    }
  }

  widget {
    alert_value_definition {
      alert_id = "895605"
      precision = 3
      unit = "b"
      text_align = "center"
      title = "Widget Title"
    }
  }

  widget {
    change_definition {
      request {
        q = "avg:system.load.1{env:staging} by {account}"
        change_type = "absolute"
        compare_to = "week_before"
        increase_good = true
        order_by = "name"
        order_dir = "desc"
        show_present = true
      }
      title = "Widget Title"
      time = {
        live_span = "1h"
      }
    }
  }

  widget {
    distribution_definition {
      request {
        q = "avg:system.load.1{env:staging} by {account}"
        style {
          palette = "warm"
        }
      }
      title = "Widget Title"
      time = {
        live_span = "1h"
      }
    }
  }

  widget {
    check_status_definition {
      check = "aws.ecs.agent_connected"
      grouping = "cluster"
      group_by = ["account", "cluster"]
      tags = ["account:demo", "cluster:awseb-ruthebdog-env-8-dn3m6u3gvk"]
      title = "Widget Title"
      time = {
        live_span = "1h"
      }
    }
  }

  widget {
    heatmap_definition {
      request {
        q = "avg:system.load.1{env:staging} by {account}"
        style {
          palette = "warm"
        }
      }
      yaxis {
        min = 1
        max = 2
        include_zero = true
        scale = "sqrt"
      }
      title = "Widget Title"
      time = {
        live_span = "1h"
      }
    }
  }

  widget {
    hostmap_definition {
      request {
        fill {
          q = "avg:system.load.1{*} by {host}"
        }
        size {
          q = "avg:memcache.uptime{*} by {host}"
        }
      }
      node_type= "container"
      group = ["host", "region"]
      no_group_hosts = true
      no_metric_hosts = true
      scope = ["region:us-east-1", "aws_account:727006795293"]
      style {
        palette = "yellow_to_green"
        palette_flip = true
        fill_min = "10"
        fill_max = "20"
      }
      title = "Widget Title"
    }
  }

  widget {
    note_definition {
      content = "note text"
      background_color = "pink"
      font_size = "14"
      text_align = "center"
      show_tick = true
      tick_edge = "left"
      tick_pos = "50%"
    }
  }

  widget {
    query_value_definition {
      request {
        q = "avg:system.load.1{env:staging} by {account}"
        aggregator = "sum"
        conditional_formats {
          comparator = "<"
          value = "2"
          palette = "white_on_green"
        }
        conditional_formats {
          comparator = ">"
          value = "2.2"
          palette = "white_on_red"
        }
      }
      autoscale = true
      custom_unit = "xx"
      precision = "4"
      text_align = "right"
      title = "Widget Title"
      time = {
        live_span = "1h"
      }
    }
  }

  widget {
    query_table_definition {
      request {
        q = "avg:system.load.1{env:staging} by {account}"
        aggregator = "sum"
        limit = "10"
        conditional_formats {
          comparator = "<"
          value = "2"
          palette = "white_on_green"
        }
        conditional_formats {
          comparator = ">"
          value = "2.2"
          palette = "white_on_red"
        }
      }
      title = "Widget Title"
      time = {
        live_span = "1h"
      }
    }
  }

  widget {
    scatterplot_definition {
      request {
        x {
          q = "avg:system.cpu.user{*} by {service, account}"
          aggregator = "max"
        }
        y {
          q = "avg:system.mem.used{*} by {service, account}"
          aggregator = "min"
        }
      }
      color_by_groups = ["account", "apm-role-group"]
      xaxis {
        include_zero = true
        label = "x"
        min = "1"
        max = "2000"
        scale = "pow"
      }
      yaxis {
        include_zero = false
        label = "y"
        min = "5"
        max = "2222"
        scale = "log"
      }
      title = "Widget Title"
      time = {
        live_span = "1h"
      }
    }
  }

  widget {
    servicemap_definition {
      service = "master-db"
      filters = ["env:prod","datacenter:us1.prod.dog"]
      title = "env: prod, datacenter:us1.prod.dog, service: master-db"
      title_size = "16"
      title_align = "left"
    }
    layout = {
      height = 43
      width = 32
      x = 5
      y = 5
    }
  }

  widget {
    timeseries_definition {
      request {
        q= "avg:system.cpu.user{app:general} by {env}"
        display_type = "line"
        style {
          palette = "warm"
          line_type = "dashed"
          line_width = "thin"
        }
        metadata {
          expression = "avg:system.cpu.user{app:general} by {env}"
          alias_name = "Alpha"
        }
      }
      request {
        log_query {
          index = "mcnulty"
          compute = {
            aggregation = "avg"
            facet = "@duration"
            interval = 5000
          }
          search = {
            query = "status:info"
          }
          group_by {
            facet = "host"
            limit = 10
            sort = {
              aggregation = "avg"
              order = "desc"
              facet = "@duration"
            }
          }
        }
        display_type = "area"
      }
      request {
        apm_query {
          index = "apm-search"
          compute = {
            aggregation = "avg"
            facet = "@duration"
            interval = 5000
          }
          search = {
            query = "type:web"
          }
          group_by {
            facet = "resource_name"
            limit = 50
            sort = {
              aggregation = "avg"
              order = "desc"
              facet = "@string_query.interval"
            }
          }
        }
        display_type = "bars"
      }
      request {
        process_query {
          metric = "process.stat.cpu.total_pct"
          search_by = "error"
          filter_by = ["active"]
          limit = 50
        }
        display_type = "area"
      }
      marker {
        display_type = "error dashed"
        label = " z=6 "
        value = "y = 4"
      }
      marker {
        display_type = "ok solid"
        value = "10 < y < 999"
        label = " x=8 "
      }
      title = "Widget Title"
      show_legend = true
      legend_size = "2"
      time = {
        live_span = "1h"
      }
      event {
        q = "sources:test tags:1"
      }
      event {
        q = "sources:test tags:2"
      }
      yaxis {
        scale = "log"
        include_zero = false
        max = 100
      }
    }
  }

  widget {
    toplist_definition {
      request {
        q= "avg:system.cpu.user{app:general} by {env}"
        conditional_formats {
          comparator = "<"
          value = "2"
          palette = "white_on_green"
        }
        conditional_formats {
          comparator = ">"
          value = "2.2"
          palette = "white_on_red"
        }
      }
      title = "Widget Title"
    }
  }

  widget {
    group_definition {
      layout_type = "ordered"
      title = "Group Widget"

      widget {
        note_definition {
          content = "cluster note widget"
          background_color = "pink"
          font_size = "14"
          text_align = "center"
          show_tick = true
          tick_edge = "left"
          tick_pos = "50%"
        }
      }

      widget {
        alert_graph_definition {
          alert_id = "123"
          viz_type = "toplist"
          title = "Alert Graph"
          time = {
            live_span = "1h"
          }
        }
      }
    }
  }

  widget {
    service_level_objective_definition {
      title = "Widget Title"
      view_type = "detail"
      slo_id = "56789"
      show_error_budget = true
      view_mode = "overall"
      time_windows = ["7d", "previous_week"]
    }
  }

  template_variable {
    name   = "var_1"
    prefix = "host"
    default = "aws"
  }
  template_variable {
    name   = "var_2"
    prefix = "service_name"
    default = "autoscaling"
  }

  template_variable_preset {
    name = "preset_1"
    template_variable {
      name = "var_1"
      value = "host.dc"
    }
    template_variable {
      name = "var_2"
      value = "my_service"
    }
  }
}

resource "datadog_dashboard" "free_dashboard_example" {
  title         = "Free Layout Dashboard"
  description   = "Created using the Datadog provider in Terraform"
  layout_type   = "free"
  is_read_only  = false

  widget {
    event_stream_definition {
      query = "*"
      event_size = "l"
      title = "Widget Title"
      title_size = 16
      title_align = "left"
      time = {
        live_span = "1h"
      }
    }
    layout = {
      height = 43
      width = 32
      x = 5
      y = 5
    }
  }

  widget {
    event_timeline_definition {
      query = "*"
      title = "Widget Title"
      title_size = 16
      title_align = "left"
      time = {
        live_span = "1h"
      }
    }
    layout = {
      height = 9
      width = 65
      x = 42
      y = 73
    }
  }

  widget {
    free_text_definition {
      text = "free text content"
      color = "#d00"
      font_size = "88"
      text_align = "left"
    }
    layout = {
      height = 20
      width = 30
      x = 42
      y = 5
    }
  }

  widget {
    iframe_definition {
      url = "http://google.com"
    }
    layout = {
      height = 46
      width = 39
      x = 111
      y = 8
    }
  }

  widget {
    image_definition {
      url = "https://images.pexels.com/photos/67636/rose-blue-flower-rose-blooms-67636.jpeg?auto=compress&cs=tinysrgb&h=350"
      sizing = "fit"
      margin = "small"
    }
    layout = {
      height = 20
      width = 30
      x = 77
      y = 7
    }
  }

  widget {
    log_stream_definition {
      indexes = ["main"]
      query = "error"
      columns = ["core_host", "core_service", "tag_source"]
      show_date_column = true
      show_message_column = true
      message_display = "expanded-md"
      sort {
        column = "time"
        order = "desc"
      }
    }
    layout = {
      height = 36
      width = 32
      x = 5
      y = 51
    }
  }

  widget {
    manage_status_definition {
      color_preference = "text"
      display_format = "countsAndList"
      hide_zero_counts = true
      query = "type:metric"
      show_last_triggered = false
      sort = "status,asc"
      summary_type = "monitors"
      title = "Widget Title"
      title_size = 16
      title_align = "left"
    }
    layout = {
      height = 40
      width = 30
      x = 112
      y = 55
    }
  }

  widget {
    trace_service_definition {
      display_format = "three_column"
      env = "datad0g.com"
      service = "alerting-cassandra"
      show_breakdown = true
      show_distribution = true
      show_errors = true
      show_hits = true
      show_latency = false
      show_resource_list = false
      size_format = "large"
      span_name = "cassandra.query"
      title = "alerting-cassandra #env:datad0g.com"
      title_align = "center"
      title_size = "13"
      time = {
        live_span = "1h"
      }
    }
    layout = {
      height = 38
      width = 67
      x = 40
      y = 28
    }
  }

  template_variable {
    name   = "var_1"
    prefix = "host"
    default = "aws"
  }
  template_variable {
    name   = "var_2"
    prefix = "service_name"
    default = "autoscaling"
  }

  template_variable_preset {
    name = "preset_1"
    template_variable {
      name = "var_1"
      value = "host.dc"
    }
    template_variable {
      name = "var_2"
      value = "my_service"
    }
  }
}
