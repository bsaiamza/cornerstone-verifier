import { useEffect, useState } from 'react'
import axios from 'axios'
import MaterialTable from '@material-table/core'
import { Refresh as RefreshIcon } from '@mui/icons-material'
import { toast } from 'react-toastify'
// components
import TypographyComponent from '../../components/Typography'
// utils
import { IAMZA_VERIFIER_URL } from '../../utils'

const apiURL = IAMZA_VERIFIER_URL + '/connections'

const ConnectionsPage = () => {
  const [data, setData] = useState([])

  useEffect(() => {
    axios
      .get(apiURL)
      .then(response => {
        setData(response.data)
      })
      .catch(error => console.log(error))
  }, [])

  const refreshConnections = async () => {
    await toast.promise(
      axios
        .get(apiURL)
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
              Connections
            </TypographyComponent>
          }
          data={data}
          columns={[
            {
              title: (
                <TypographyComponent variant="h6">Name</TypographyComponent>
              ),
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
          ]}
          actions={[
            {
              icon: () => <RefreshIcon />,
              tooltip: 'Refresh connections',
              isFreeAction: true,
              onClick: () => refreshConnections(),
            },
          ]}
          options={{
            actionsColumnIndex: -1,
          }}
        />
      </div>
      <div style={{ marginBottom: '2rem' }} />
    </>
  )
}

export default ConnectionsPage
