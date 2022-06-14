import MaterialTable from '@material-table/core'
import { Refresh } from '@mui/icons-material'
import axios from 'axios'
import { useEffect, useState } from 'react'
import { toast } from 'react-toastify'
// components
import TypographyComponent from '../../components/Typography'

const PresentationRecordsPage = () => {
  const [data, setData] = useState([])

  useEffect(() => {
    let apiURL = '/api/v1/cornerstone/verifier/presentations'

    if (process.env.API_BASE_URL) {
      apiURL = process.env.API_BASE_URL + '/cornerstone/verifier/presentations'
    }

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      axios
        .get(process.env.REACT_APP_API_URL + 'presentations')
        .then(response => {
          console.log(response.data)
          setData(response.data)
        })
        .catch(error => console.log(error))
    } else {
      axios
        .get(apiURL)
        .then(response => {
          console.log(response.data)
          setData(response.data)
        })
        .catch(error => console.log(error))
    }
  }, [])

  const refreshVerificationRecords = async () => {
    let apiURL = '/api/v1/cornerstone/verifier/presentations'

    if (process.env.API_BASE_URL) {
      apiURL = process.env.API_BASE_URL + '/cornerstone/verifier/presentations'
    }

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      await toast.promise(
        axios
          .get(process.env.REACT_APP_API_URL + 'presentations')
          .then(response => {
            setData(response.data)
            toast.success('Refreshed verification records!')
          })
          .catch(error => console.log(error)),
        {
          pending: 'Refreshing...',
        }
      )
    } else {
      await toast.promise(
        axios
          .get(apiURL)
          .then(response => {
            setData(response.data)
            toast.success('Refreshed verification records!')
          })
          .catch(error => console.log(error)),
        {
          pending: 'Refreshing...',
        }
      )
    }
  }

  return (
    <>
      <div style={{ margin: '2rem' }}>
        <MaterialTable
          style={{ padding: '1rem' }}
          title={
            <TypographyComponent
              variant="h5"
              sx={{ textDecoration: 'underline' }}
            >
              Verification Records
            </TypographyComponent>
          }
          data={data}
          columns={[
            {
              title: (
                <TypographyComponent variant="h6">
                  Created On
                </TypographyComponent>
              ),
              field: 'created_at',
              type: 'datetime',
            },
            {
              title: (
                <TypographyComponent variant="h6">
                  Connection ID
                </TypographyComponent>
              ),
              field: 'connection_id',
            },
            {
              title: (
                <TypographyComponent variant="h6">
                  Presentation Exchange ID
                </TypographyComponent>
              ),
              field: 'presentation_exchange_id',
            },
            {
              title: (
                <TypographyComponent variant="h6">
                  Verification State
                </TypographyComponent>
              ),
              field: 'state',
            },
            {
              title: (
                <TypographyComponent variant="h6">Verified</TypographyComponent>
              ),
              field: 'verified',
            },
            {
              title: (
                <TypographyComponent variant="h6">
                  Updated On
                </TypographyComponent>
              ),
              field: 'updated_at',
              type: 'datetime',
            },
          ]}
          actions={[
            {
              icon: () => <Refresh />,
              tooltip: 'Refresh records',
              isFreeAction: true,
              onClick: () => refreshVerificationRecords(),
            },
            // rowData => ({
            //   icon: () => <CredentialIcon />,
            //   tooltip: 'Issue Credential',
            //   onClick: () => {
            //     handleOpen()
            //     setConnectionId(rowData.connection_id)
            //   },
            // }),
          ]}
          options={{
            actionsColumnIndex: -1,
          }}
        />
        {/* {renderModal} */}
      </div>
      <div style={{ marginBottom: '2rem' }} />
    </>
  )
}

export default PresentationRecordsPage
