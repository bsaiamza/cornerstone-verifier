import { ListItemText } from '@mui/material'
import React from 'react'

const ListItemTextComponent = ({ color, primary, sx }) => {
  return <ListItemText color={color} primary={primary} sx={{ sx }} />
}

export default ListItemTextComponent
