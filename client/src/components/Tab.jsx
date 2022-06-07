import { Tab } from '@mui/material'

const TabComponent = ({ label, ...props }) => {
  return <Tab label={label} {...props} />
}

export default TabComponent
