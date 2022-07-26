import { Tooltip } from '@mui/material'

const TooltipComponent = ({ children, title }) => {
  return <Tooltip title={title}>{children}</Tooltip>
}

export default TooltipComponent
