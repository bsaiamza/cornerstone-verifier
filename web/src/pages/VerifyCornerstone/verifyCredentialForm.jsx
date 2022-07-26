import { useState } from 'react'
import {
  Divider,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
} from '@mui/material'
import { toast } from 'react-toastify'
import axios from 'axios'
import QRCode from 'react-qr-code'
// components
import ButtonComponent from '../../components/Button'
import FormikComponent from '../../components/Formik'
import FormComponent from '../../components/Form'
import TypographyComponent from '../../components/Typography'
// utils
import { IAMZA_VERIFIER_URL } from '../../utils'

const apiURL = IAMZA_VERIFIER_URL + '/verify-cornerstone'

const VerifyCredentialForm = () => {
  const [submitting, setSubmitting] = useState(false)
  const [success, setSuccess] = useState(false)
  const [data, setData] = useState([])

  const sendOffer = async data => {
    setSubmitting(true)

    await toast.promise(
      axios
        .post(apiURL, data)
        .then(response => {
          setData(response.data)
          setSuccess(true)
          toast.success('Verification request generated!')
        })
        .catch(error => {
          toast.error(error.response.data.msg)
        }),
      {
        pending: 'Generating request...',
      }
    )

    setSubmitting(false)
  }

  return (
    <>
      <TypographyComponent variant="h5">
        Verify my Cornerstone Credential
      </TypographyComponent>
      <Divider />
      <div style={{ marginTop: '1rem' }}>
        <FormikComponent
          initialValues={{
            id_number: false,
            forenames: false,
            surname: false,
            gender: false,
            date_of_birth: false,
            country_of_birth: false,
          }}
          onSubmit={(values, { resetForm }) => {
            sendOffer(values)
            // resetForm()
          }}
        >
          {({ values, handleChange, touched, errors }) => (
            <FormComponent>
              <div>
                <FormControl sx={{ width: '16.5rem' }}>
                  <InputLabel id="id_number" sx={{ margin: '1rem 0 0 1rem' }}>
                    ID Number
                  </InputLabel>
                  <Select
                    labelId="id_number"
                    id="id_number"
                    name="id_number"
                    value={values.id_number}
                    label="ID Number"
                    onChange={handleChange}
                    sx={{ m: '1rem' }}
                  >
                    <MenuItem value={false}>False</MenuItem>
                    <MenuItem value={true}>True</MenuItem>
                  </Select>
                </FormControl>
              </div>
              <div>
                <FormControl sx={{ width: '16.5rem' }}>
                  <InputLabel id="forenames" sx={{ margin: '1rem 0 0 1rem' }}>
                    Forenames
                  </InputLabel>
                  <Select
                    labelId="forenames"
                    id="forenames"
                    name="forenames"
                    value={values.forenames}
                    label="Forenames"
                    onChange={handleChange}
                    sx={{ m: '1rem' }}
                  >
                    <MenuItem value={false}>False</MenuItem>
                    <MenuItem value={true}>True</MenuItem>
                  </Select>
                </FormControl>
              </div>
              <div>
                <FormControl sx={{ width: '16.5rem' }}>
                  <InputLabel id="surname" sx={{ margin: '1rem 0 0 1rem' }}>
                    Surname
                  </InputLabel>
                  <Select
                    labelId="surname"
                    id="surname"
                    name="surname"
                    value={values.surname}
                    label="surname"
                    onChange={handleChange}
                    sx={{ m: '1rem' }}
                  >
                    <MenuItem value={false}>False</MenuItem>
                    <MenuItem value={true}>True</MenuItem>
                  </Select>
                </FormControl>
              </div>
              <div>
                <FormControl sx={{ width: '16.5rem' }}>
                  <InputLabel id="gender" sx={{ margin: '1rem 0 0 1rem' }}>
                    Gender
                  </InputLabel>
                  <Select
                    labelId="gender"
                    id="gender"
                    name="gender"
                    value={values.gender}
                    label="gender"
                    onChange={handleChange}
                    sx={{ m: '1rem' }}
                  >
                    <MenuItem value={false}>False</MenuItem>
                    <MenuItem value={true}>True</MenuItem>
                  </Select>
                </FormControl>
              </div>
              <div>
                <FormControl sx={{ width: '16.5rem' }}>
                  <InputLabel
                    id="date_of_birth"
                    sx={{ margin: '1rem 0 0 1rem' }}
                  >
                    D.O.B
                  </InputLabel>
                  <Select
                    labelId="date_of_birth"
                    id="date_of_birth"
                    name="date_of_birth"
                    value={values.date_of_birth}
                    label="date_of_birth"
                    onChange={handleChange}
                    sx={{ m: '1rem' }}
                  >
                    <MenuItem value={false}>False</MenuItem>
                    <MenuItem value={true}>True</MenuItem>
                  </Select>
                </FormControl>
              </div>
              <div>
                <FormControl sx={{ width: '16.5rem' }}>
                  <InputLabel
                    id="country_of_birth"
                    sx={{ margin: '1rem 0 0 1rem' }}
                  >
                    Country of Birth
                  </InputLabel>
                  <Select
                    labelId="country_of_birth"
                    id="country_of_birth"
                    name="country_of_birth"
                    value={values.country_of_birth}
                    label="country_of_birth"
                    onChange={handleChange}
                    sx={{ m: '1rem' }}
                  >
                    <MenuItem value={false}>False</MenuItem>
                    <MenuItem value={true}>True</MenuItem>
                  </Select>
                </FormControl>
              </div>

              <div>
                <ButtonComponent
                  variant="contained"
                  type="submit"
                  sx={{ color: '#fff', m: '1rem' }}
                  disabled={submitting}
                >
                  Submit
                </ButtonComponent>
              </div>
            </FormComponent>
          )}
        </FormikComponent>
      </div>

      {success && data.proofRequest ? <QRCode value={data.proofRequest} /> : ''}
    </>
  )
}

export default VerifyCredentialForm
