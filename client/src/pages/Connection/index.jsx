import { useEffect, useState } from 'react'
import axios from 'axios'
import MaterialTable from '@material-table/core'
import {
  Badge as CredentialIcon,
  Refresh as RefreshIcon,
} from '@mui/icons-material'
import { toast } from 'react-toastify'
// components
import TypographyComponent from '../../components/Typography'
import CredentialModal from './modal'

const ConnectionsPage = () => {
  const [data, setData] = useState([])
  const [open, setOpen] = useState(false)
  const [connectionId, setConnectionId] = useState('')

  const handleOpen = () => setOpen(true)
  const handleClose = () => setOpen(false)

  useEffect(() => {
    let apiURL = '/api/v1/cornerstone/issuer/connections'

    if (process.env.API_BASE_URL) {
      apiURL = process.env.API_BASE_URL + '/cornerstone/issuer/connections'
    }

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      axios
        .get(process.env.REACT_APP_API_URL + 'connections?state=active')
        .then(response => {
          setData(response.data)
        })
        .catch(error => console.log(error))
    } else {
      axios
        .get(apiURL + '?state=active')
        .then(response => {
          setData(response.data)
        })
        .catch(error => console.log(error))
    }
  }, [])

  const refreshConnections = async () => {
    let apiURL = '/api/v1/cornerstone/issuer/connections'

    if (process.env.API_BASE_URL) {
      apiURL = process.env.API_BASE_URL + '/cornerstone/issuer/connections'
    }

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      await toast.promise(
        axios
          .get(process.env.REACT_APP_API_URL + 'connections?state=active')
          .then(response => {
            setData(response.data)
            toast.success('Refreshed connections!')
          })
          .catch(error => console.log(error)),
        {
          pending: 'Refreshing...',
        }
      )
    } else {
      await toast.promise(
        axios
          .get(apiURL + '?state=active')
          .then(response => {
            setData(response.data)
            toast.success('Refreshed connections!')
          })
          .catch(error => console.log(error)),
        {
          pending: 'Refreshing...',
        }
      )
    }
  }

  const renderModal = (
    <CredentialModal
      connectionId={connectionId}
      handleClose={handleClose}
      open={open}
    />
  )

  return (
    <div style={{ margin: '2rem' }}>
      <MaterialTable
        style={{ padding: '1rem' }}
        title={
          <TypographyComponent
            variant="h5"
            sx={{ textDecoration: 'underline' }}
          >
            Connections
          </TypographyComponent>
        }
        data={data}
        columns={[
          {
            title: <TypographyComponent variant="h6">Name</TypographyComponent>,
            field: 'their_label',
          },
          {
            title: (
              <TypographyComponent variant="h6">
                Connected On
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
                Connection State
              </TypographyComponent>
            ),
            field: 'state',
          },
        ]}
        actions={[
          {
            icon: () => <RefreshIcon />,
            tooltip: 'Refresh connections',
            isFreeAction: true,
            onClick: () => refreshConnections(),
          },
          rowData => ({
            icon: () => <CredentialIcon />,
            tooltip: 'Issue Credential',
            onClick: () => {
              handleOpen()
              setConnectionId(rowData.connection_id)
            },
          }),
        ]}
        options={{
          actionsColumnIndex: -1,
        }}
      />
      {renderModal}
    </div>
  )
}

export default ConnectionsPage
