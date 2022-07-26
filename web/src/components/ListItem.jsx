import { ListItem } from '@mui/material'

const ListItemComponent = ({ children, disablePadding, onClick, sx }) => {
  return (
    <ListItem disablePadding={disablePadding} onClick={onClick} sx={sx}>
      {children}
    </ListItem>
  )
}

export default ListItemComponent
