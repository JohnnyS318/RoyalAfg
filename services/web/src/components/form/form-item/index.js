import React from "react";
import PropTypes from "prop-types";

const FormItem = (props) => {
    return <div className="mb-6 font-sans text-lg font-medium">{props.children}</div>;
};

FormItem.propTypes = {
    children: PropTypes.oneOfType([PropTypes.element, PropTypes.array])
};

export default FormItem;
