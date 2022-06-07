import { useState } from 'react'
// components
import BoxComponent from '../../components/Box'
import TabsComponent from '../../components/Tabs'
import TabComponent from '../../components/Tab'
import { a11yProps, TabPanel } from '../../components/TabPanel'
// verify
import PresentationRecords from './records'
import NewProofRequestForm from './form'

const VerifyPage = () => {
  const [value, setValue] = useState(0)

  const handleChange = (event, newValue) => {
    setValue(newValue)
  }

  return (
    <div style={{ margin: '1rem' }}>
      <BoxComponent>
        <TabsComponent
          value={value}
          onChange={handleChange}
          ariaLabel="Verify Tabs"
        >
          <TabComponent label="All" {...a11yProps(0)} />
          <TabComponent label="New" {...a11yProps(1)} />
        </TabsComponent>
      </BoxComponent>
      <TabPanel value={value} index={0}>
        <PresentationRecords />
      </TabPanel>
      <TabPanel value={value} index={1}>
        <NewProofRequestForm />
      </TabPanel>
    </div>
  )
}

export default VerifyPage
