// mui5
import { AppBar } from '@mui/material'

const AppBarComponent = ({ children, position, sx }) => {
  return (
    <AppBar position={position} sx={sx}>
      {children}
    </AppBar>
  )
}

export default AppBarComponent
