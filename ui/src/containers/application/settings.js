import React, { useState } from 'react';
import useForm from 'react-hook-form';
import { useNavigation } from 'react-navi';
import * as yup from 'yup';
import { toaster } from 'evergreen-ui';

import api from '../../api';
import utils from '../../utils';
import validators from '../../validators';
import Card from '../../components/card';
import Popup from '../../components/popup';
import Field from '../../components/field';
import Alert from '../../components/alert';
import { Button, Text, Form } from '../../components/core';

const validationSchema = yup.object().shape({
  name: validators.name.required(),
  description: yup.string(),
});

const ApplicationSettings = ({
  route: {
    data: { params, application },
  },
}) => {
  const { register, handleSubmit, errors, formState } = useForm({
    validationSchema,
    defaultValues: {
      name: application.name,
      description: application.description,
    },
  });
  const [backendError, setBackendError] = useState();
  const [showDeletePopup, setShowDeletePopup] = useState();
  const navigation = useNavigation();

  const submit = async data => {
    try {
      await api.updateApplication({
        projectId: params.project,
        applicationId: application.id,
        data,
      });
      toaster.success('Application updated successfully.');
      navigation.navigate(`/${params.project}/applications/${data.name}`);
    } catch (error) {
      setBackendError(utils.parseError(error));
      toaster.danger('Application was not updated.');
      console.log(error);
    }
  };

  const submitDelete = async () => {
    setBackendError(null);
    try {
      await api.deleteApplication({
        projectId: params.project,
        applicationId: application.name,
      });
      toaster.success('Successfully deleted application.');
      navigation.navigate(`/${params.project}/applications`);
    } catch (error) {
      setBackendError(utils.parseError(error));
      toaster.danger('Application was not deleted.');
      console.log(error);
    }
    setShowDeletePopup(false);
  };

  return (
    <>
      <Card
        title="Application Settings"
        size="medium"
        actions={[
          {
            title: 'Delete',
            onClick: () => setShowDeletePopup(true),
            variant: 'secondary',
          },
        ]}
      >
        <Alert show={backendError} variant="error" description={backendError} />
        <Form
          onSubmit={e => {
            setBackendError(null);
            handleSubmit(submit)(e);
          }}
        >
          <Field
            autoFocus
            required
            label="Name"
            name="name"
            ref={register}
            errors={errors.name}
          />
          <Field
            type="textarea"
            label="Description"
            name="description"
            ref={register}
            errors={errors.description}
          />
          <Button title="Update" type="submit" disabled={!formState.dirty} />
        </Form>
      </Card>
      <Popup show={showDeletePopup} onClose={() => setShowDeletePopup(false)}>
        <Card title="Delete Application" border size="medium">
          <Text>
            You are about to delete the <strong>{application.name}</strong>{' '}
            application.
          </Text>
          <Button marginTop={5} title="Delete" onClick={submitDelete} />
        </Card>
      </Popup>
    </>
  );
};

export default ApplicationSettings;
