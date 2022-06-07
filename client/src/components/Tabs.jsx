import { Tabs } from '@mui/material'

const TabsComponent = ({ ariaLabel, children, onChange, value }) => {
  return (
    <Tabs aria-label={ariaLabel} onChange={onChange} value={value}>
      {children}
    </Tabs>
  )
}

export default TabsComponent
