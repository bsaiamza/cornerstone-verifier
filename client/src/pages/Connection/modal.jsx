import { useState } from 'react'
import { Avatar, Divider, Modal } from '@mui/material'
import { Badge as CredentialIcon } from '@mui/icons-material'
import { toast } from 'react-toastify'
import axios from 'axios'
import QRCode from 'react-qr-code'
import validate from 'za-id-validator'
// components
import BoxComponent from '../../components/Box'
import ButtonComponent from '../../components/Button'
import FormikComponent from '../../components/Formik'
import FormComponent from '../../components/Form'
import TextFieldComponent from '../../components/TextField'
import TypographyComponent from '../../components/Typography'

const idValidation = data => {
  const errors = {}

  if (validate(data.id_number)) {
  } else {
    errors.id_number = 'Invalid ZA ID number!'
  }

  return errors
}

const CredentialModal = ({ connectionId, handleClose, open }) => {
  const [submitting, setSubmitting] = useState(false)
  const [success, setSuccess] = useState(false)
  const [data, setData] = useState({})

  const sendOffer = async data => {
    setSubmitting(true)
    let apiURL = '/api/v1/cornerstone/issuer/credential'

    if (process.env.API_BASE_URL) {
      apiURL = process.env.API_BASE_URL + '/cornerstone/issuer/credential'
    }

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      await toast.promise(
        axios
          .post(process.env.REACT_APP_API_URL + 'credential', data)
          .then(response => {
            setData(response.data)
            setSuccess(true)
            toast.success('Credential qr code generated!')
          })
          .catch(error => {
            toast.error(error.response.data.msg)
          }),
        {
          pending: 'Generating credential qr code...',
        }
      )
    } else {
      await toast.promise(
        axios
          .post(apiURL, data)
          .then(response => {
            setData(response.data)
            setSuccess(true)
            toast.success('Credential qr code generated!')
          })
          .catch(error => {
            toast.error(error.response.data.msg)
          }),
        {
          pending: 'Generating credential qr code...',
        }
      )
    }
    setSubmitting(false)
  }

  return (
    <Modal onClose={handleClose} open={open}>
      <BoxComponent
        sx={{
          position: 'absolute',
          top: '50%',
          left: '50%',
          transform: 'translate(-50%, -50%)',
          boxShadow: 24,
          backgroundColor: 'background.paper',
          borderRadius: '1rem',
          p: '3rem 4rem 3rem 3rem',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'center',
          textAlign: 'center',
        }}
      >
        <Avatar sx={{ m: 1 }}>
          <CredentialIcon />
        </Avatar>
        <TypographyComponent variant="h5">
          Issue a Cornerstone Credential
        </TypographyComponent>
        <Divider />
        <FormikComponent
          initialValues={{
            email: '',
            id_number: '',
            surname: '',
            forenames: '',
            gender: '',
            date_of_birth: '',
            country_of_birth: '',
          }}
          validate={idValidation}
          onSubmit={(values, { resetForm }) => {
            sendOffer(values)
            // resetForm()
          }}
        >
          {({ values, handleChange, touched, errors }) => (
            <FormComponent>
              <div>
                <TextFieldComponent
                  error={touched.id_number && Boolean(errors.id_number)}
                  helperText={touched.id_number && errors.id_number}
                  id="id_number"
                  name="id_number"
                  value={values.id_number}
                  onChange={handleChange}
                  label="ID Number"
                  sx={{ m: '1rem' }}
                  required
                />

                <TextFieldComponent
                  id="forenames"
                  name="forenames"
                  value={values.forenames}
                  onChange={handleChange}
                  label="Forenames"
                  sx={{ m: '1rem' }}
                  required
                />
              </div>

              <div>
                <TextFieldComponent
                  id="surname"
                  name="surname"
                  value={values.surname}
                  onChange={handleChange}
                  label="Surname"
                  sx={{ m: '1rem' }}
                  required
                />
                <TextFieldComponent
                  id="gender"
                  name="gender"
                  // value={values.gender}
                  value={
                    (values.gender =
                      values.id_number.substring(6, 7) > 4 ? 'Male' : 'Female')
                  }
                  onChange={handleChange}
                  label="Gender"
                  sx={{ m: '1rem' }}
                  required
                  disabled
                />
              </div>

              <div>
                <TextFieldComponent
                  id="date_of_birth"
                  name="date_of_birth"
                  // type="date"
                  // value={values.date_of_birth}
                  value={
                    (values.date_of_birth =
                      values.id_number.substring(0, 1) > 2
                        ? '19' +
                          values.id_number.substring(0, 2) +
                          '-' +
                          values.id_number.substring(2, 4) +
                          '-' +
                          values.id_number.substring(4, 6)
                        : '20' +
                          values.id_number.substring(0, 2) +
                          '-' +
                          values.id_number.substring(2, 4) +
                          '-' +
                          values.id_number.substring(4, 6))
                  }
                  onChange={handleChange}
                  label="D.O.B"
                  // sx={{ m: '1rem', width: '15rem' }}
                  sx={{ m: '1rem' }}
                  required
                  // InputLabelProps={{
                  //   shrink: true,
                  // }}
                  disabled
                />

                <TextFieldComponent
                  id="country_of_birth"
                  name="country_of_birth"
                  value={
                    (values.country_of_birth =
                      values.id_number.substring(10, 11) === '0'
                        ? 'South Africa'
                        : '')
                  }
                  onChange={handleChange}
                  label="Country of Birth"
                  sx={{ m: '1rem' }}
                  required
                  disabled={values.id_number.substring(10, 11) === '0'}
                />
              </div>

              <div>
                <TextFieldComponent
                  id="email"
                  name="email"
                  value={values.email}
                  onChange={handleChange}
                  label="Email"
                  sx={{ m: '1rem' }}
                  required
                />
              </div>

              <div>
                <ButtonComponent
                  variant="contained"
                  type="submit"
                  sx={{ color: '#fff', m: '1rem' }}
                  disabled={submitting}
                >
                  Send
                </ButtonComponent>
              </div>
            </FormComponent>
          )}
        </FormikComponent>

        {success ? <QRCode value={data.credential} /> : ''}
      </BoxComponent>
    </Modal>
  )
}

export default CredentialModal
