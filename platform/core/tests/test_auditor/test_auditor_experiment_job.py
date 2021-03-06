# pylint:disable=ungrouped-imports

from unittest.mock import patch

import pytest

import auditor

from events.registry import experiment_job as experiment_job_events
from factories.factory_experiments import ExperimentJobFactory
from tests.test_auditor.utils import AuditorBaseTest


@pytest.mark.auditor_mark
class AuditorExperimentJobTest(AuditorBaseTest):
    """Testing subscribed events"""
    EVENTS = experiment_job_events.EVENTS

    def setUp(self):
        super().setUp()
        self.experiment_job = ExperimentJobFactory()
        self.tested_events = {
            experiment_job_events.EXPERIMENT_JOB_VIEWED,
            experiment_job_events.EXPERIMENT_JOB_RESOURCES_VIEWED,
            experiment_job_events.EXPERIMENT_JOB_LOGS_VIEWED,
            experiment_job_events.EXPERIMENT_JOB_STATUSES_VIEWED,
            experiment_job_events.EXPERIMENT_JOB_NEW_STATUS,
        }

    @patch('executor.executor_service.ExecutorService.record_event')
    @patch('notifier.service.NotifierService.record_event')
    @patch('tracker.service.TrackerService.record_event')
    @patch('activitylogs.service.ActivityLogService.record_event')
    def test_experiment_job_viewed(self,
                                   activitylogs_record,
                                   tracker_record,
                                   notifier_record,
                                   executor_record):
        auditor.record(event_type=experiment_job_events.EXPERIMENT_JOB_VIEWED,
                       instance=self.experiment_job,
                       actor_id=1,
                       actor_name='foo')

        assert tracker_record.call_count == 1
        assert activitylogs_record.call_count == 1
        assert notifier_record.call_count == 0
        assert executor_record.call_count == 0

    @patch('executor.executor_service.ExecutorService.record_event')
    @patch('notifier.service.NotifierService.record_event')
    @patch('tracker.service.TrackerService.record_event')
    @patch('activitylogs.service.ActivityLogService.record_event')
    def test_experiment_resources_viewed(self,
                                         activitylogs_record,
                                         tracker_record,
                                         notifier_record,
                                         executor_record):
        auditor.record(event_type=experiment_job_events.EXPERIMENT_JOB_RESOURCES_VIEWED,
                       instance=self.experiment_job,
                       actor_id=1,
                       actor_name='foo')

        assert tracker_record.call_count == 1
        assert activitylogs_record.call_count == 1
        assert notifier_record.call_count == 0
        assert executor_record.call_count == 0

    @patch('executor.executor_service.ExecutorService.record_event')
    @patch('notifier.service.NotifierService.record_event')
    @patch('tracker.service.TrackerService.record_event')
    @patch('activitylogs.service.ActivityLogService.record_event')
    def test_experiment_logs_viewed(self,
                                    activitylogs_record,
                                    tracker_record,
                                    notifier_record,
                                    executor_record):
        auditor.record(event_type=experiment_job_events.EXPERIMENT_JOB_LOGS_VIEWED,
                       instance=self.experiment_job,
                       actor_id=1,
                       actor_name='foo')

        assert tracker_record.call_count == 1
        assert activitylogs_record.call_count == 1
        assert notifier_record.call_count == 0
        assert executor_record.call_count == 0

    @patch('executor.executor_service.ExecutorService.record_event')
    @patch('notifier.service.NotifierService.record_event')
    @patch('tracker.service.TrackerService.record_event')
    @patch('activitylogs.service.ActivityLogService.record_event')
    def test_experiment_job_statuses_viewed(self,
                                            activitylogs_record,
                                            tracker_record,
                                            notifier_record,
                                            executor_record):
        auditor.record(event_type=experiment_job_events.EXPERIMENT_JOB_STATUSES_VIEWED,
                       instance=self.experiment_job,
                       actor_id=1,
                       actor_name='foo')

        assert tracker_record.call_count == 1
        assert activitylogs_record.call_count == 1
        assert notifier_record.call_count == 0
        assert executor_record.call_count == 0

    @patch('executor.executor_service.ExecutorService.record_event')
    @patch('notifier.service.NotifierService.record_event')
    @patch('tracker.service.TrackerService.record_event')
    @patch('activitylogs.service.ActivityLogService.record_event')
    def test_experiment_job_new_status(self,
                                       activitylogs_record,
                                       tracker_record,
                                       notifier_record,
                                       executor_record):
        auditor.record(event_type=experiment_job_events.EXPERIMENT_JOB_NEW_STATUS,
                       instance=self.experiment_job)

        assert tracker_record.call_count == 0
        assert activitylogs_record.call_count == 0
        assert notifier_record.call_count == 0
        assert executor_record.call_count == 1


del AuditorBaseTest
