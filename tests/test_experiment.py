# -*- coding: utf-8 -*-
from __future__ import absolute_import, division, print_function

import uuid
from unittest import TestCase

from datetime import datetime

from polyaxon_schemas.experiment import (
    ExperimentConfig,
    JobLabelConfig,
    PodStateConfig,
    JobStateConfig,
    ContainerGPUResourcesConfig,
    ContainerResourcesConfig,
    ExperimentJobConfig,
    ExperimentStatusConfig,
    ExperimentJobStatusConfig)


class TestExperimentConfigs(TestCase):
    def test_experiment_config(self):
        config_dict = {
            'uuid': uuid.uuid4().hex,
            'project': uuid.uuid4().hex,
            'project_name': 'name.name',
            'group': uuid.uuid4().hex,
            'group_name': 'name.name.1',
            'unique_name': 'user.proj.1',
            'last_status': 'Running',
            'description': 'description',
            'content': 'content',
            'config': {'k': 'v'},
            'num_jobs': 1,
        }
        config = ExperimentConfig.from_dict(config_dict)
        assert config.to_dict() == config_dict

        config_dict.pop('description')
        config_dict.pop('content')
        config_dict.pop('config')
        config_dict.pop('project')
        config_dict.pop('group')
        assert config.to_light_dict() == config_dict

    def test_experiment_with_jobs_config(self):
        config_dict = {'sequence': 2,
                       'config': {},
                       'content': '',
                       'unique_name': 'adam.proj.1',
                       'uuid': uuid.uuid4().hex,
                       'project': uuid.uuid4().hex,
                       'project_name': 'user.name',
                       'group': uuid.uuid4().hex,
                       'group_name': 'user.name.1',
                       'last_status': 'Running',
                       'num_jobs': 1,
                       'jobs': [ExperimentJobConfig(uuid=uuid.uuid4().hex,
                                                    experiment=uuid.uuid4().hex,
                                                    experiment_name='name.name.1',
                                                    created_at=datetime.now(),
                                                    updated_at=datetime.now(),
                                                    definition={}).to_dict()]}
        config = ExperimentConfig.from_dict(config_dict)
        assert config.to_dict() == config_dict

    def test_experiment_job_config(self):
        config_dict = {'uuid': uuid.uuid4().hex,
                       'experiment': uuid.uuid4().hex,
                       'experiment_name': 'name.name',
                       'created_at': datetime.now().isoformat(),
                       'updated_at': datetime.now().isoformat(),
                       'definition': {},
                       'role': 'master',
                       'sequence': 1,
                       'unique_name': 'project.1.1.master'}
        config = ExperimentJobConfig.from_dict(config_dict)
        config_to_dict = config.to_dict()
        config_to_dict.pop('created_at')
        config_to_dict.pop('updated_at')
        config_dict.pop('created_at')
        config_dict.pop('updated_at')
        assert config_to_dict == config_dict

        config_dict.pop('definition')
        config_to_dict = config.to_light_dict()
        config_to_dict.pop('created_at')
        config_to_dict.pop('updated_at')
        config_dict.pop('sequence')
        config_dict.pop('unique_name')
        assert config_to_dict == config_dict

    def test_experiment_status_config(self):
        config_dict = {'uuid': uuid.uuid4().hex,
                       'experiment': uuid.uuid4().hex,
                       'created_at': datetime.now().isoformat(),
                       'status': 'Running'}
        config = ExperimentStatusConfig.from_dict(config_dict)
        config_to_dict = config.to_dict()
        config_to_dict.pop('created_at')
        config_dict.pop('created_at')
        assert config_to_dict == config_dict

    def test_experiment_status_config(self):
        config_dict = {'uuid': uuid.uuid4().hex,
                       'job': uuid.uuid4().hex,
                       'created_at': datetime.now().isoformat(),
                       'status': 'Running'}
        config = ExperimentJobStatusConfig.from_dict(config_dict)
        config_to_dict = config.to_dict()
        config_to_dict.pop('created_at')
        config_dict.pop('created_at')
        assert config_to_dict == config_dict

        config_to_dict = config.to_light_dict()
        config_to_dict.pop('created_at')
        config_dict.pop('details', '')
        assert config_to_dict == config_dict

    @staticmethod
    def create_pod_labels():
        project_name = 'user/test-id1'
        experiment_group_name = 'user/test-id1/1'
        experiment_name = 'user/test-id1/1/2'
        project_uuid = uuid.uuid4().hex
        experiment_group_uuid = uuid.uuid4().hex
        experiment_uuid = uuid.uuid4().hex
        job_uuid = uuid.uuid4().hex
        task_type = 'master'
        return {'project_name': project_name,
                'experiment_group_name': experiment_group_name,
                'experiment_name': experiment_name,
                'project_uuid': project_uuid,
                'experiment_group_uuid': experiment_group_uuid,
                'experiment_uuid': experiment_uuid,
                'task_type': task_type,
                'task_idx': '0',
                'job_uuid': job_uuid,
                'role': 'polyaxon-worker',
                'type': 'polyaxon-experiment'
                }

    @classmethod
    def create_pod_state(cls):
        event_type = 'ADDED'
        phase = 'Running'
        labels = cls.create_pod_labels()
        deletion_timestamp = datetime.now().isoformat()
        pod_conditions = {}
        container_statuses = {}
        return {
            'event_type': event_type,
            'phase': phase,
            'labels': labels,
            'deletion_timestamp': deletion_timestamp,
            'pod_conditions': pod_conditions,
            'container_statuses': container_statuses,
        }

    def test_pod_label_config(self):
        config_dict = self.create_pod_labels()
        config = JobLabelConfig.from_dict(config_dict)
        assert config.to_dict() == config_dict

    def test_pod_state_config(self):
        config_dict = self.create_pod_state()
        config = PodStateConfig.from_dict(config_dict)
        config_to_dict = config.to_dict()
        config_to_dict.pop('deletion_timestamp')
        config_dict.pop('deletion_timestamp')
        assert config_to_dict == config_dict

    def test_job_state_config(self):
        config_dict = {
            'status': 'Running',
            'message': 'something',
            'details': self.create_pod_state()
        }
        config = JobStateConfig.from_dict(config_dict)
        config_to_dict = config.to_dict()
        config_to_dict['details'].pop('deletion_timestamp')
        config_dict['details'].pop('deletion_timestamp')
        assert config_to_dict == config_dict

    def test_container_gpu_resources(self):
        config_dict = {
            'index': 0,
            'bus_id': '0000:00:1E.1',
            'memory_free': 1000,
            'memory_total': 12883853312,
            'memory_used': 8388608000,
            'memory_utilization': 0,
            'minor': 1,
            'name': 'GeForce GTX TITAN 0',
            'power_draw': 125,
            'power_limit': 250,
            'processes': [{'command': 'python',
                           'gpu_memory_usage': 4000,
                           'pid': 48448,
                           'username': 'user1'},
                          {'command': 'python',
                           'gpu_memory_usage': 4000,
                           'pid': 153223,
                           'username': 'user2'}],
            'serial': '0322917092147',
            'temperature_gpu': 80,
            'utilization_gpu': 76,
            'uuid': 'GPU-10fb0fbd-2696-43f3-467f-d280d906a107'
        }
        config = ContainerGPUResourcesConfig.from_dict(config_dict)
        config_to_dict = config.to_dict()
        assert config_to_dict == config_dict

    def test_container_resources(self):
        gpu_resources = {
            'index': 0,
            'bus_id': '0000:00:1E.1',
            'memory_free': 1000,
            'memory_total': 12883853312,
            'memory_used': 8388608000,
            'memory_utilization': 0,
            'minor': 1,
            'name': 'GeForce GTX TITAN 0',
            'power_draw': 125,
            'power_limit': 250,
            'processes': [{'command': 'python',
                           'gpu_memory_usage': 4000,
                           'pid': 48448,
                           'username': 'user1'},
                          {'command': 'python',
                           'gpu_memory_usage': 4000,
                           'pid': 153223,
                           'username': 'user2'}],
            'serial': '0322917092147',
            'temperature_gpu': 80,
            'utilization_gpu': 76,
            'uuid': 'GPU-10fb0fbd-2696-43f3-467f-d280d906a107'
        }

        config_dict = {
            'job_uuid': uuid.uuid4().hex,
            'experiment_uuid': uuid.uuid4().hex,
            'container_id': '3175e88873af9077688cee20eaadc0c07746efb84d01ae696d6d17ed9bcdfbc4',
            'cpu_percentage': 0.6947691836734693,
            'percpu_percentage': [0.4564075715616173, 0.23836161211185192],
            'memory_used': 84467712,
            'memory_limit': 2096160768,
            'gpu_resources': gpu_resources
        }
        config = ContainerResourcesConfig.from_dict(config_dict)
        config_to_dict = config.to_dict()
        assert config_to_dict == config_dict
