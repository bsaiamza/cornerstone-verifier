import { Formik } from 'formik'

const FormikComponent = ({ children, initialValues, onSubmit, validate }) => {
  return (
    <Formik
      initialValues={initialValues}
      onSubmit={onSubmit}
      validate={validate}
    >
      {children}
    </Formik>
  )
}

export default FormikComponent
