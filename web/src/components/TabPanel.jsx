import PropTypes from 'prop-types'
// components
import BoxComponent from './Box'
import TypographyComponent from './Typography'

export const TabPanel = props => {
  const { children, value, index, ...other } = props

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`custom-tabpanel-${index}`}
      aria-labelledby={`custom-tab-${index}`}
      {...other}
    >
      {value === index && (
        <BoxComponent sx={{ p: 3 }}>
          <TypographyComponent>{children}</TypographyComponent>
        </BoxComponent>
      )}
    </div>
  )
}

TabPanel.propTypes = {
  children: PropTypes.node,
  index: PropTypes.number.isRequired,
  value: PropTypes.number.isRequired,
}

export const a11yProps = index => {
  return {
    id: `custom-tab-${index}`,
    'aria-controls': `custom-tabpanel-${index}`,
  }
}
