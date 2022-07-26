import { useState } from 'react'
import { Divider } from '@mui/material'
import { toast } from 'react-toastify'
import axios from 'axios'
// components
import ButtonComponent from '../../components/Button'
import FormikComponent from '../../components/Formik'
import FormComponent from '../../components/Form'
import TextFieldComponent from '../../components/TextField'
import TypographyComponent from '../../components/Typography'
// utils
import { IAMZA_VERIFIER_URL } from '../../utils'

const apiURL = IAMZA_VERIFIER_URL + '/verify-email'

const VerifyCredentialEmailForm = () => {
  const [submitting, setSubmitting] = useState(false)

  const sendOffer = async data => {
    setSubmitting(true)

    await toast.promise(
      axios
        .post(apiURL, data)
        .then(response => {
          toast.success('Verification request emailed!')
        })
        .catch(error => {
          toast.error(error.response.data.msg)
        }),
      {
        pending: 'Emailing request...',
      }
    )

    setSubmitting(false)
  }

  return (
    <>
      <TypographyComponent variant="h5">
        Verify a Credential
      </TypographyComponent>

      <Divider />

      <TypographyComponent variant="h6">
        The verification request is based off the following IAMZA credentials:
      </TypographyComponent>

      <div>
        <a
          href="https://issuer.iamza-sandbox.com/"
          target="_blank"
          rel="noreferrer"
          style={{ textDecoration: 'none', color: '#faa61a' }}
        >
          Cornerstone
        </a>
      </div>

      <div>
        <a
          href="https://addrissuer.iamza-sandbox.com/"
          target="_blank"
          rel="noreferrer"
          style={{ textDecoration: 'none', color: '#faa61a' }}
        >
          Address
        </a>
      </div>

      <div>
        <a
          href="https://vacissuer.iamza-sandbox.com/"
          target="_blank"
          rel="noreferrer"
          style={{ textDecoration: 'none', color: '#faa61a' }}
        >
          Vaccine
        </a>
      </div>

      <Divider />
      <div style={{ marginTop: '1rem' }}>
        <FormikComponent
          initialValues={{
            email: '',
          }}
          onSubmit={(values, { resetForm }) => {
            sendOffer(values)
            // resetForm()
          }}
        >
          {({ values, handleChange, touched, errors }) => (
            <FormComponent>
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
                  Submit
                </ButtonComponent>
              </div>
            </FormComponent>
          )}
        </FormikComponent>
      </div>
    </>
  )
}

export default VerifyCredentialEmailForm
