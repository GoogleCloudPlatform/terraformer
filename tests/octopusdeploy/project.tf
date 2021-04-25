resource "octopusdeploy_project" "deploymark_api" {
  name                  = "Deploymark API"
  description           = "Annotate deployments"
  lifecycle_id          = octopusdeploy_lifecycle.stage_prod.id
  project_group_id      = octopusdeploy_project_group.utilities.id
  skip_machine_behavior = "SkipUnavailableMachines"

  deployment_step {
    windows_service {
      executable_path                = "batch_processor\\batch_processor_service.exe"
      service_name                   = "Billing Batch Processor"
      step_name                      = "Deploy Billing Batch Processor Windows Service"
      step_condition                 = "failure"
      package                        = "Billing.BatchProcessor"
      json_file_variable_replacement = "appsettings.json"

      target_roles = [
        "Billing-Batch-Processor",
      ]
    }

    inline_script {
      step_name   = "Cleanup Temporary Files"
      script_type = "PowerShell"

      script_body = <<EOF
        $oldFiles = Get-ChildItem -Path 'C:\billing_archived_jobs'
        Remove-Item $oldFiles -Force -Recurse
        EOF

      target_roles = [
        "Billing-Batch-Processor",
      ]
    }
  }
}
