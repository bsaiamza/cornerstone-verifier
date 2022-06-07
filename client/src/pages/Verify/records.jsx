import axios from 'axios'
import { useEffect, useState } from 'react'
// components
import CardComponent from '../../components/Card'
import GridComponent from '../../components/Grid'
import ListItemComponent from '../../components/ListItem'
import ListItemTextComponent from '../../components/ListItemText'

const PresentationRecords = () => {
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

  return (
    <GridComponent container justify="center" spacing={2}>
      {data ? (
        data.length === 0 ? (
          <CardComponent
            sx={{
              m: '1rem',
            }}
          >
            No records available!
          </CardComponent>
        ) : (
          data.map((request, index) => (
            <GridComponent item xs={12} md={4}>
              <CardComponent
                key={index}
                sx={{
                  m: '1rem',
                  wordWrap: 'break-word',
                }}
              >
                <ListItemComponent>
                  <ListItemTextComponent
                    primary={'Created: ' + request.created_at}
                  />
                </ListItemComponent>
                <ListItemComponent>
                  <ListItemTextComponent
                    primary={'Connection ID: ' + request.connection_id}
                  />
                </ListItemComponent>
                <ListItemComponent>
                  <ListItemTextComponent
                    primary={'Presentation Exchange ID: ' + request.pres_ex_id}
                  />
                </ListItemComponent>

                {request.error_msg.length !== 0 ? (
                  <ListItemComponent>
                    <ListItemTextComponent primary={request.error_msg} />
                  </ListItemComponent>
                ) : (
                  ''
                )}

                <ListItemComponent>
                  <ListItemTextComponent primary={'State: ' + request.state} />
                </ListItemComponent>

                <ListItemComponent>
                  <ListItemTextComponent
                    primary={'Verified: ' + request.verified}
                  />
                </ListItemComponent>

                <ListItemComponent sx={{ backgroundColor: '#eee' }}>
                  <ListItemTextComponent
                    primary={'Last updated at: ' + request.updated_at}
                  />
                </ListItemComponent>
              </CardComponent>
            </GridComponent>
          ))
        )
      ) : (
        <CardComponent
          sx={{
            m: '1rem',
          }}
        >
          No records available!
        </CardComponent>
      )}
    </GridComponent>
  )
}

export default PresentationRecords
