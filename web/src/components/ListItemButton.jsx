import { ListItemButton } from '@mui/material'

const ListItemButtonComponent = ({ children, onClick, sx }) => {
  return <ListItemButton sx={sx}>{children}</ListItemButton>
}

export default ListItemButtonComponent
