import React from 'react';
import useForm from 'react-hook-form';
import { useNavigation } from 'react-navi';

import Card from '../../components/card';
import Field from '../../components/field';
import { Button, Form } from '../../components/core';
import api from '../../api';

const DeviceRegistrationTokenSettings = () => {
  const navigation = useNavigation();
  const { register, handleSubmit } = useForm();

  const submit = data => {
    api.updateDeviceRegistrationToken(data).then(() => {
      navigation.navigate(`/${projectId}/provisioning`);
    });
  };

  return (
    <Card title="Device Registration Token Settings">
      <Form onSubmit={handleSubmit(submit)}>
        <Field label="Name" name="name" ref={register} />
        <Field
          type="textarea"
          label="Description"
          name="description"
          ref={register}
        />
        <Field
          label="Maximum Device Registrations"
          name="maxRegistrations"
          description="Limit the number of devices that can be registered using this token"
          hint="Leave empty to allow unlimited registrations"
          ref={register}
        />
        <Button title="Update Settings" />
        <Button
          title="Delete Device Registration Token"
          variant="tertiary"
          onClick={() => this.setState({ showDeleteDialog: true })}
        />
        {/* <Dialog
          isShown={this.state.showDeleteDialog}
          title="Delete Device Registration Token"
          intent="danger"
          onCloseComplete={() => this.setState({ showDeleteDialog: false })}
          onConfirm={() => this.handleDelete()}
          confirmLabel="Delete Device Registration Token"
        >
          You are about to delete the{" "}
          <strong>{this.props.deviceRegistrationToken.name}</strong> Device
          Registration Token.
        </Dialog> */}
      </Form>
    </Card>
  );
};

// export default DeviceRegistrationTokenSettings;

// export default class DeviceRegistrationTokenSettings extends Component {
//   state = {
//     name: this.props.deviceRegistrationToken.name,
//     description: this.props.deviceRegistrationToken.description,
//     maxRegistrations: (typeof this.props.deviceRegistrationToken.maxRegistrations === 'number' ?
//       String(this.props.deviceRegistrationToken.maxRegistrations) :
//       ''),
//     nameValidationMessage: null,
//     maxRegistrationsValidationMessage: null,
//     unchanged: true,
//     showDeleteDialog: false,
//     backendError: null
//   };

//   handleUpdate = () => {
// var nameValidationMessage = utils.checkName('Device Registration Token', this.state.name);

// //always set validation message for name
// this.setState({
//   nameValidationMessage: nameValidationMessage,
//   backendError: null
// });

// if (nameValidationMessage !== null) {
//   return;
// }

// // Convert max registrations to int or undefined
// var maxRegistrationsCleaned;
// if (this.state.maxRegistrations === '') {
//   maxRegistrationsCleaned = null;
// } else {
//   maxRegistrationsCleaned = Number(this.state.maxRegistrations);

//   if (isNaN(maxRegistrationsCleaned)) {
//     this.setState({
//       maxRegistrationsValidationMessage: 'Max Registrations should either be a number or be left empty.'
//     });
//     return;
//   }
// }

//   {
//     withCredentials: true
//   }
// )
// .then(response => {
//   toaster.success('Device Registration Token updated successfully.');

// })
// .catch(error => {
//   if (utils.is4xx(error.response.status)) {
//     this.setState({
//       backendError: utils.convertErrorMessage(error.response.data)
//     });
//   } else {
//     toaster.danger('Device Registration Token was not updated.');
//     console.log(error);
//   }
// });
//};

// handleDelete() {

//         withCredentials: true
//       }
//     )
//     .then(response => {
//       this.setState({
//         showDeleteDialog: false
//       });
//       toaster.success('Successfully deleted Device Registration Token.');
//       this.props.history.push(
//         `/${this.props.projectName}/provisioning`
//       );
//     })
//     .catch(error => {
//       this.setState({
//         showDeleteDialog: false
//       });
//       if (utils.is4xx(error.response.status)) {
//         this.setState({
//           backendError: utils.convertErrorMessage(error.response.data)
//         });
//       } else {
//         toaster.danger('Device Registration Token was not deleted.');
//         console.log(error);
//       }
//     });
// }