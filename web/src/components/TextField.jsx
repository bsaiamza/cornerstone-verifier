import { TextField } from '@mui/material'

const FieldComponent = ({
  defaultValue,
  disabled,
  error,
  focused,
  fullWidth,
  helperText,
  id,
  inputProps,
  InputLabelProps,
  label,
  name,
  onChange,
  props,
  required,
  sx,
  type,
  value,
}) => {
  return (
    <TextField
      defaultValue={defaultValue}
      disabled={disabled}
      error={error}
      focused={focused}
      fullWidth={fullWidth}
      helperText={helperText}
      id={id}
      inputProps={inputProps}
      InputLabelProps={InputLabelProps}
      label={label}
      name={name}
      onChange={onChange}
      required={required}
      sx={sx}
      type={type}
      value={value}
      {...props}
    />
  )
}

export default FieldComponent
