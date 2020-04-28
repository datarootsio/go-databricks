package jobs

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "encoding/json"
    "github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/innovationnorway/go-databricks/jobs"

        // LifeCycleState enumerates the values for life cycle state.
    type LifeCycleState string

    const (
                // INTERNALERROR ...
        INTERNALERROR LifeCycleState = "INTERNAL_ERROR"
                // PENDING ...
        PENDING LifeCycleState = "PENDING"
                // RUNNING ...
        RUNNING LifeCycleState = "RUNNING"
                // SKIPPED ...
        SKIPPED LifeCycleState = "SKIPPED"
                // TERMINATED ...
        TERMINATED LifeCycleState = "TERMINATED"
                // TERMINATING ...
        TERMINATING LifeCycleState = "TERMINATING"
            )
    // PossibleLifeCycleStateValues returns an array of possible values for the LifeCycleState const type.
    func PossibleLifeCycleStateValues() []LifeCycleState {
        return []LifeCycleState{INTERNALERROR,PENDING,RUNNING,SKIPPED,TERMINATED,TERMINATING}
    }

        // ResultState enumerates the values for result state.
    type ResultState string

    const (
                // CANCELED ...
        CANCELED ResultState = "CANCELED"
                // FAILED ...
        FAILED ResultState = "FAILED"
                // SUCCESS ...
        SUCCESS ResultState = "SUCCESS"
                // TIMEDOUT ...
        TIMEDOUT ResultState = "TIMEDOUT"
            )
    // PossibleResultStateValues returns an array of possible values for the ResultState const type.
    func PossibleResultStateValues() []ResultState {
        return []ResultState{CANCELED,FAILED,SUCCESS,TIMEDOUT}
    }

        // Trigger enumerates the values for trigger.
    type Trigger string

    const (
                // ONETIME ...
        ONETIME Trigger = "ONE_TIME"
                // PERIODIC ...
        PERIODIC Trigger = "PERIODIC"
                // RETRY ...
        RETRY Trigger = "RETRY"
            )
    // PossibleTriggerValues returns an array of possible values for the Trigger const type.
    func PossibleTriggerValues() []Trigger {
        return []Trigger{ONETIME,PERIODIC,RETRY}
    }

        // Type enumerates the values for type.
    type Type string

    const (
                // DASHBOARD ...
        DASHBOARD Type = "DASHBOARD"
                // NOTEBOOK ...
        NOTEBOOK Type = "NOTEBOOK"
            )
    // PossibleTypeValues returns an array of possible values for the Type const type.
    func PossibleTypeValues() []Type {
        return []Type{DASHBOARD,NOTEBOOK}
    }

            // Attributes ...
            type Attributes struct {
            autorest.Response `json:"-"`
            JobID *string `json:"job_id,omitempty"`
            }

            // ClusterInstance ...
            type ClusterInstance struct {
            ClusterID *string `json:"cluster_id,omitempty"`
            SparkContextID *string `json:"spark_context_id,omitempty"`
            }

            // ClusterSpec ...
            type ClusterSpec struct {
            ExistingClusterID *string `json:"existing_cluster_id,omitempty"`
            NewCluster interface{} `json:"new_cluster,omitempty"`
            Libraries *[]interface{} `json:"libraries,omitempty"`
            }

            // CreateAttributes ...
            type CreateAttributes struct {
            ExistingClusterID *string `json:"existing_cluster_id,omitempty"`
            NewCluster interface{} `json:"new_cluster,omitempty"`
            NotebookTask *NotebookTask `json:"notebook_task,omitempty"`
            SparkJarTask *SparkJarTask `json:"spark_jar_task,omitempty"`
            SparkPythonTask *SparkPythonTask `json:"spark_python_task,omitempty"`
            SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty"`
            Name *string `json:"name,omitempty"`
            Libraries *[]interface{} `json:"libraries,omitempty"`
            EmailNotification *EmailNotifications `json:"email_notification,omitempty"`
            TimeoutSeconds *int32 `json:"timeout_seconds,omitempty"`
            MaxRetries *int32 `json:"max_retries,omitempty"`
            MinRetryIntervalMillis *int32 `json:"min_retry_interval_millis,omitempty"`
            RetryOnTimeout *bool `json:"retry_on_timeout,omitempty"`
            Schedule *CronSchedule `json:"schedule,omitempty"`
            MaxConcurrentRuns *int32 `json:"max_concurrent_runs,omitempty"`
            }

            // CronSchedule ...
            type CronSchedule struct {
            QuartzCronExpression *string `json:"quartz_cron_expression,omitempty"`
            TimezoneID *string `json:"timezone_id,omitempty"`
            }

            // EmailNotifications ...
            type EmailNotifications struct {
            OnStart *[]string `json:"on_start,omitempty"`
            OnSuccess *[]string `json:"on_success,omitempty"`
            OnFailure *[]string `json:"on_failure,omitempty"`
            NoAlertForSkippedRuns *bool `json:"no_alert_for_skipped_runs,omitempty"`
            }

            // GetResult ...
            type GetResult struct {
            autorest.Response `json:"-"`
            JobID *int64 `json:"job_id,omitempty"`
            CreatorUserName *string `json:"creator_user_name,omitempty"`
            Settings *CreateAttributes `json:"settings,omitempty"`
            CreatedTime *int64 `json:"created_time,omitempty"`
            }

            // Job ...
            type Job struct {
            JobID *int64 `json:"job_id,omitempty"`
            CreatorUserName *string `json:"creator_user_name,omitempty"`
            Settings *CreateAttributes `json:"settings,omitempty"`
            CreatedTime *int64 `json:"created_time,omitempty"`
            }

            // ListResult ...
            type ListResult struct {
            autorest.Response `json:"-"`
            JobsProperty *[]Job `json:"jobs,omitempty"`
            }

            // NotebookOutput ...
            type NotebookOutput struct {
            Result *string `json:"result,omitempty"`
            Truncated *bool `json:"truncated,omitempty"`
            }

            // NotebookTask ...
            type NotebookTask struct {
            NotebookPath *string `json:"notebook_path,omitempty"`
            BaseParameters map[string]map[string]*string `json:"base_parameters"`
            }

        // MarshalJSON is the custom marshaler for NotebookTask.
        func (nt NotebookTask)MarshalJSON() ([]byte, error){
        objectMap := make(map[string]interface{})
                if(nt.NotebookPath != nil) {
                objectMap["notebook_path"] = nt.NotebookPath
                }
                if(nt.BaseParameters != nil) {
                objectMap["base_parameters"] = nt.BaseParameters
                }
                return json.Marshal(objectMap)
        }

            // ResetAttributes ...
            type ResetAttributes struct {
            JobID *int64 `json:"job_id,omitempty"`
            NewSettings *CreateAttributes `json:"new_settings,omitempty"`
            }

            // Run ...
            type Run struct {
            JobID *int64 `json:"job_id,omitempty"`
            RunID *int64 `json:"run_id,omitempty"`
            CreatorUserName *string `json:"creator_user_name,omitempty"`
            NumberInJob *int64 `json:"number_in_job,omitempty"`
            OriginalAttemptRunID *int64 `json:"original_attempt_run_id,omitempty"`
            State *RunState `json:"state,omitempty"`
            Schedule *CronSchedule `json:"schedule,omitempty"`
            Task *Task `json:"task,omitempty"`
            ClusterSpec *ClusterSpec `json:"cluster_spec,omitempty"`
            ClusterInstance *ClusterInstance `json:"cluster_instance,omitempty"`
            OverridingParameters *RunParameters `json:"overriding_parameters,omitempty"`
            StartTime *int64 `json:"start_time,omitempty"`
            SetupDuration *int64 `json:"setup_duration,omitempty"`
            ExecutionDuration *int64 `json:"execution_duration,omitempty"`
            CleanupDuration *int64 `json:"cleanup_duration,omitempty"`
            // Trigger - Possible values include: 'PERIODIC', 'ONETIME', 'RETRY'
            Trigger Trigger `json:"trigger,omitempty"`
            }

            // RunParameters ...
            type RunParameters struct {
            JarParams *[]string `json:"jar_params,omitempty"`
            NotebookParams map[string]map[string]*string `json:"notebook_params"`
            PythonParams *[]string `json:"python_params,omitempty"`
            SparkSubmitParams *[]string `json:"spark_submit_params,omitempty"`
            }

        // MarshalJSON is the custom marshaler for RunParameters.
        func (rp RunParameters)MarshalJSON() ([]byte, error){
        objectMap := make(map[string]interface{})
                if(rp.JarParams != nil) {
                objectMap["jar_params"] = rp.JarParams
                }
                if(rp.NotebookParams != nil) {
                objectMap["notebook_params"] = rp.NotebookParams
                }
                if(rp.PythonParams != nil) {
                objectMap["python_params"] = rp.PythonParams
                }
                if(rp.SparkSubmitParams != nil) {
                objectMap["spark_submit_params"] = rp.SparkSubmitParams
                }
                return json.Marshal(objectMap)
        }

            // RunState ...
            type RunState struct {
            // LifeCycleState - Possible values include: 'PENDING', 'RUNNING', 'TERMINATING', 'TERMINATED', 'SKIPPED', 'INTERNALERROR'
            LifeCycleState LifeCycleState `json:"life_cycle_state,omitempty"`
            // ResultState - Possible values include: 'SUCCESS', 'FAILED', 'TIMEDOUT', 'CANCELED'
            ResultState ResultState `json:"result_state,omitempty"`
            StateMessage *string `json:"state_message,omitempty"`
            }

            // SparkJarTask ...
            type SparkJarTask struct {
            JarURI *string `json:"jar_uri,omitempty"`
            MainClassName *string `json:"main_class_name,omitempty"`
            Parameters *[]string `json:"parameters,omitempty"`
            }

            // SparkPythonTask ...
            type SparkPythonTask struct {
            PythonFile *string `json:"python_file,omitempty"`
            Parameters *[]string `json:"parameters,omitempty"`
            }

            // SparkSubmitTask ...
            type SparkSubmitTask struct {
            Parameters *[]string `json:"parameters,omitempty"`
            }

            // Task ...
            type Task struct {
            NotebookTask *NotebookTask `json:"notebook_task,omitempty"`
            SparkJarTask *SparkJarTask `json:"spark_jar_task,omitempty"`
            SparkPythonTask *SparkPythonTask `json:"spark_python_task,omitempty"`
            SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty"`
            }

            // ViewItem ...
            type ViewItem struct {
            Content *string `json:"content,omitempty"`
            Name *string `json:"name,omitempty"`
            // Type - Possible values include: 'NOTEBOOK', 'DASHBOARD'
            Type Type `json:"type,omitempty"`
            }

