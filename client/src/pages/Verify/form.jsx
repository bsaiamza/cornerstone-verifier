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
import CardComponent from '../../components/Card'
import FormikComponent from '../../components/Formik'
import FormComponent from '../../components/Form'
import GridComponent from '../../components/Grid'
import TypographyComponent from '../../components/Typography'

const NewProofRequestForm = () => {
  const [submitting, setSubmitting] = useState(false)
  const [success, setSuccess] = useState(false)
  const [data, setData] = useState([])

  const handleSubmit = async values => {
    setSubmitting(true)
    let apiURL = '/api/v1/cornerstone/verifier/proof'

    if (process.env.API_BASE_URL) {
      apiURL = process.env.API_BASE_URL + '/cornerstone/verifier/proof'
    }

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      await toast.promise(
        axios
          .post(process.env.REACT_APP_API_URL + 'proof', values)
          .then(response => {
            setData(response.data)
            setSuccess(true)
            toast.success('Sent proof request!')
          })
          .catch(error => {
            toast.error(error.response.data.msg)
          }),
        {
          pending: 'Sending...',
        }
      )
    } else {
      await toast.promise(
        axios
          .post(apiURL, values)
          .then(response => {
            setData(response.data)
            setSuccess(true)
            toast.success('Sent proof request!')
          })
          .catch(error => {
            toast.error(error.response.data.msg)
          }),
        {
          pending: 'Sending...',
        }
      )
    }
    setSubmitting(false)
  }

  return (
    <GridComponent container justify="center" spacing={2}>
      <GridComponent item xs={12} md={6}>
        <CardComponent
          sx={{
            m: '1rem',
          }}
        >
          <TypographyComponent variant="h6">
            New Cornerstone Credential Proof Request
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
                handleSubmit(values)
                resetForm()
              }}
            >
              {({ values, handleChange }) => (
                <FormComponent>
                  <div>
                    <FormControl sx={{ width: '16.5rem' }}>
                      <InputLabel
                        id="id_number"
                        sx={{ margin: '1rem 0 0 1rem' }}
                      >
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
                      <InputLabel
                        id="forenames"
                        sx={{ margin: '1rem 0 0 1rem' }}
                      >
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
                      Send
                    </ButtonComponent>
                  </div>
                </FormComponent>
              )}
            </FormikComponent>
          </div>
          {success ? <QRCode value={data.proofRequest} /> : ''}
        </CardComponent>
      </GridComponent>
    </GridComponent>
  )
}

export default NewProofRequestForm
