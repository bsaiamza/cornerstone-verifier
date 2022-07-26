import { useEffect, useState } from 'react'
import { Divider, Tooltip } from '@mui/material'
import { Refresh as RefreshIcon } from '@mui/icons-material'
import axios from 'axios'
import QRCode from 'react-qr-code'
// components
import ButtonComponent from '../../components/Button'
import BoxComponent from '../../components/Box'
import TypographyComponent from '../../components/Typography'
// images
import scan from '../../assets/images/scan.png'
// utils
import { IAMZA_VERIFIER_URL } from '../../utils'
import { toast } from 'react-toastify'

const apiURL = IAMZA_VERIFIER_URL + '/verify'

const VerifyCredentialEmailForm = () => {
  const [url, setURL] = useState('')

  useEffect(() => {
    axios
      .get(apiURL)
      .then(response => {
        console.log(response.data)
        setURL(response.data.proofRequest)
      })
      .catch(error => console.log(error))
  }, [])

  const refreshRequest = async () => {
    await toast.promise(
      axios
        .get(apiURL)
        .then(response => {
          setURL(response.data.proofRequest)
          toast.success('Refreshed verification request!')
        })
        .catch(error => console.log(error)),
      {
        pending: 'Refreshing...',
      }
    )
  }

  return (
    <>
      <TypographyComponent variant="h5">
        Verify a Credential
      </TypographyComponent>

      <Divider />

      <TypographyComponent variant="h6">
        The below verification request is based off the following IAMZA
        credentials:
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

      <BoxComponent
        alt="DI"
        component="img"
        src={scan}
        sx={{
          height: '10rem',
          width: '10rem',
          opacity: 0.5,
        }}
      />
      <div>
        <QRCode value={url} />
      </div>

      <div>
        <ButtonComponent
          variant="outlined"
          onClick={refreshRequest}
          sx={{ marginTop: '2rem' }}
        >
          <Tooltip title="Refresh" arrow>
            <RefreshIcon />
          </Tooltip>
        </ButtonComponent>
      </div>
    </>
  )
}

export default VerifyCredentialEmailForm
