import { Input } from '@mui/material'

const LabelComponent = ({ defaultValue, disabled }) => {
  return <Input disabled={disabled} defaultValue={defaultValue} />
}

export default LabelComponent
