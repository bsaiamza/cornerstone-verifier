import { Card, CardContent } from '@mui/material'

const CardComponent = ({ children, sx }) => {
  return (
    <Card sx={sx} elevation={3}>
      <CardContent>{children}</CardContent>
    </Card>
  )
}

export default CardComponent
