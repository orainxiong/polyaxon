#!/usr/bin/python
#
# Copyright 2019 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8
from __future__ import absolute_import, division, print_function

import uuid

from collections import OrderedDict
from unittest import TestCase

import pytest

from marshmallow import ValidationError
from tests.utils import assert_equal_dict

from polyaxon.schemas.polyflow.io import IOConfig
from polyaxon.schemas.polyflow.params import ParamSpec, get_param
from polyaxon.types import types


@pytest.mark.polyflow_mark
class TestIOConfigs(TestCase):
    def test_wrong_io_config(self):
        # No name
        with self.assertRaises(ValidationError):
            IOConfig.from_dict({})

    def test_unsupported_io_config_type(self):
        with self.assertRaises(ValidationError):
            IOConfig.from_dict({"name": "input1", "type": "something"})

    def test_wrong_io_config_default(self):
        with self.assertRaises(ValidationError):
            IOConfig.from_dict({"name": "input1", "type": types.FLOAT, "value": "foo"})

        with self.assertRaises(ValidationError):
            IOConfig.from_dict({"name": "input1", "type": types.GCS, "value": 234})

    def test_wrong_io_config_flag(self):
        with self.assertRaises(ValidationError):
            IOConfig.from_dict({"name": "input1", "type": types.S3, "is_flag": True})

        with self.assertRaises(ValidationError):
            IOConfig.from_dict({"name": "input1", "type": types.FLOAT, "is_flag": True})

    def test_io_config_optionals(self):
        config_dict = {"name": "input1"}
        config = IOConfig.from_dict(config_dict)
        assert_equal_dict(config.to_dict(), config_dict)

    def test_io_config_desc(self):
        # test desc
        config_dict = {"name": "input1", "description": "some text"}
        config = IOConfig.from_dict(config_dict)
        assert_equal_dict(config.to_dict(), config_dict)

    def test_io_config_types(self):
        config_dict = {"name": "input1", "description": "some text", "type": types.INT}
        config = IOConfig.from_dict(config_dict)
        assert_equal_dict(config.to_dict(), config_dict)
        expected_repr = OrderedDict((("name", "input1"), ("type", "int"), ("value", 3)))
        assert config.get_repr_from_value(3) == expected_repr
        assert config.get_repr() == OrderedDict((("name", "input1"), ("type", "int")))

        config_dict = {"name": "input1", "description": "some text", "type": types.S3}
        config = IOConfig.from_dict(config_dict)
        assert_equal_dict(config.to_dict(), config_dict)
        expected_repr = OrderedDict(
            (("name", "input1"), ("type", types.S3), ("value", "s3://foo"))
        )
        assert config.get_repr_from_value("s3://foo") == expected_repr
        assert config.get_repr() == OrderedDict(
            (("name", "input1"), ("type", types.S3))
        )

    def test_io_config_default(self):
        config_dict = {
            "name": "input1",
            "description": "some text",
            "type": types.BOOL,
            "is_optional": True,
            "value": True,
        }
        config = IOConfig.from_dict(config_dict)
        assert_equal_dict(config.to_dict(), config_dict)
        expected_repr = OrderedDict(
            (("name", "input1"), ("type", "bool"), ("value", True))
        )
        assert config.get_repr_from_value(None) == expected_repr
        assert config.get_repr() == expected_repr

        config_dict = {
            "name": "input1",
            "description": "some text",
            "type": types.FLOAT,
            "is_optional": True,
            "value": 3.4,
        }
        config = IOConfig.from_dict(config_dict)
        assert_equal_dict(config.to_dict(), config_dict)
        expected_repr = OrderedDict(
            (("name", "input1"), ("type", "float"), ("value", 3.4))
        )
        assert config.get_repr_from_value(None) == expected_repr
        assert config.get_repr() == expected_repr

    def test_io_config_default_and_required(self):
        config_dict = {
            "name": "input1",
            "description": "some text",
            "type": types.BOOL,
            "value": True,
            "is_optional": True,
        }
        config = IOConfig.from_dict(config_dict)
        assert_equal_dict(config.to_dict(), config_dict)

        config_dict = {
            "name": "input1",
            "description": "some text",
            "type": types.STR,
            "value": "foo",
        }
        with self.assertRaises(ValidationError):
            IOConfig.from_dict(config_dict)

    def test_io_config_required(self):
        config_dict = {
            "name": "input1",
            "description": "some text",
            "type": "float",
            "is_optional": False,
        }
        config = IOConfig.from_dict(config_dict)
        assert_equal_dict(config.to_dict(), config_dict)
        expected_repr = OrderedDict(
            (("name", "input1"), ("type", "float"), ("value", 1.1))
        )
        assert config.get_repr_from_value(1.1) == expected_repr
        assert config.get_repr() == OrderedDict((("name", "input1"), ("type", "float")))

    def test_io_config_flag(self):
        config_dict = {
            "name": "input1",
            "description": "some text",
            "type": types.BOOL,
            "is_flag": True,
        }
        config = IOConfig.from_dict(config_dict)
        assert_equal_dict(config.to_dict(), config_dict)
        expected_repr = OrderedDict(
            (("name", "input1"), ("type", "bool"), ("value", False))
        )
        assert config.get_repr_from_value(False) == expected_repr

    def test_value_non_typed_input(self):
        config_dict = {"name": "input1"}
        config = IOConfig.from_dict(config_dict)
        assert config.validate_value("foo") == "foo"
        assert config.validate_value(1) == 1
        assert config.validate_value(True) is True

        expected_repr = OrderedDict((("name", "input1"), ("value", "foo")))
        assert config.get_repr_from_value("foo") == expected_repr
        assert config.get_repr() == OrderedDict(name="input1")

    def test_value_typed_input(self):
        config_dict = {"name": "input1", "type": types.BOOL}
        config = IOConfig.from_dict(config_dict)
        with self.assertRaises(ValidationError):
            config.validate_value("foo")
        with self.assertRaises(ValidationError):
            config.validate_value(1)
        with self.assertRaises(ValidationError):
            config.validate_value(None)

        assert config.validate_value(True) is True

    def test_value_typed_input_with_default(self):
        config_dict = {
            "name": "input1",
            "type": types.INT,
            "value": 12,
            "is_optional": True,
        }
        config = IOConfig.from_dict(config_dict)
        with self.assertRaises(ValidationError):
            config.validate_value("foo")

        assert config.validate_value(1) == 1
        assert config.validate_value(0) == 0
        assert config.validate_value(-1) == -1
        assert config.validate_value(None) == 12
        expected_repr = OrderedDict(
            (("name", "input1"), ("type", "int"), ("value", 12))
        )
        assert config.get_repr_from_value(None) == expected_repr
        assert config.get_repr() == expected_repr

    def test_get_param(self):
        # None string values should exit fast
        assert get_param(
            name="foo", value=1, iotype=types.INT, is_flag=False
        ) == ParamSpec(
            name="foo",
            iotype=types.INT,
            value=1,
            entity=None,
            entity_ref=None,
            entity_value=None,
            is_flag=False,
        )

        # Str values none regex
        assert get_param(
            name="foo", value="1", iotype=types.INT, is_flag=False
        ) == ParamSpec(
            name="foo",
            iotype=types.INT,
            value="1",
            entity=None,
            entity_ref=None,
            entity_value=None,
            is_flag=False,
        )

        assert get_param(
            name="foo", value="SDfd", iotype=types.STR, is_flag=False
        ) == ParamSpec(
            name="foo",
            iotype=types.STR,
            value="SDfd",
            entity=None,
            entity_ref=None,
            entity_value=None,
            is_flag=False,
        )

        # Regex validation dag
        assert get_param(
            name="foo", value="{{ dag.inputs.foo }}", iotype=types.BOOL, is_flag=True
        ) == ParamSpec(
            name="foo",
            iotype=types.BOOL,
            value="dag.inputs.foo",
            entity="dag",
            entity_ref="_",
            entity_value="foo",
            is_flag=True,
        )

        # Regex validation dag: invalid params
        with self.assertRaises(ValidationError):
            get_param(
                name="foo",
                value="{{ dag.outputs.foo }}",
                iotype=types.BOOL,
                is_flag=True,
            )
        with self.assertRaises(ValidationError):
            get_param(
                name="foo",
                value="{{ dag.1.inputs.foo }}",
                iotype=types.BOOL,
                is_flag=True,
            )
        with self.assertRaises(ValidationError):
            get_param(
                name="foo", value="{{ dag.inputs }}", iotype=types.BOOL, is_flag=True
            )

        # Regex validation ops
        assert get_param(
            name="foo",
            value="{{ ops.foo-bar.outputs.foo }}",
            iotype=types.BOOL,
            is_flag=True,
        ) == ParamSpec(
            name="foo",
            iotype=types.BOOL,
            value="ops.foo-bar.outputs.foo",
            entity="ops",
            entity_ref="foo-bar",
            entity_value="foo",
            is_flag=True,
        )
        assert get_param(
            name="foo",
            value="{{ ops.foo-bar.inputs.foo }}",
            iotype=types.BOOL,
            is_flag=True,
        ) == ParamSpec(
            name="foo",
            iotype=types.BOOL,
            value="ops.foo-bar.inputs.foo",
            entity="ops",
            entity_ref="foo-bar",
            entity_value="foo",
            is_flag=True,
        )

        # Regex validation ops: invalid params
        with self.assertRaises(ValidationError):
            get_param(
                name="foo",
                value="{{ ops.foo-bar.outputs }}",
                iotype=types.BOOL,
                is_flag=True,
            )
        with self.assertRaises(ValidationError):
            get_param(
                name="foo",
                value="{{ ops.foo-bar.inputs }}",
                iotype=types.BOOL,
                is_flag=True,
            )

        # Regex validation runs
        uid = uuid.uuid4().hex
        assert get_param(
            name="foo",
            value="{{" + "runs.{}.outputs.foo".format(uid) + "}}",
            iotype=types.BOOL,
            is_flag=True,
        ) == ParamSpec(
            name="foo",
            iotype=types.BOOL,
            value="runs.{}.outputs.foo".format(uid),
            entity="runs",
            entity_ref=uid,
            entity_value="foo",
            is_flag=True,
        )

        # Regex validation runs: invalid params
        with self.assertRaises(ValidationError):
            get_param(
                name="foo",
                value="{{ runs.foo-bar.outputs.foo }}",
                iotype=types.BOOL,
                is_flag=True,
            )
        with self.assertRaises(ValidationError):
            get_param(
                name="foo",
                value="{{" + "runs.{}.inputs.foo".format(uid) + "}}",
                iotype=types.BOOL,
                is_flag=True,
            )
