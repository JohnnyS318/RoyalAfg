import React from "react";
import Link from "next/link";

const BackToAccount = () => {
    return (
        <Link href="/account">
            <span className="absolute cursor-pointer font-sans font-semibold text-sm ml-6 mt-4 py-1 px-3 bg-gray-300  rounded-full hover:bg-gray-800 hover:text-white transition-colors duration-200 ease-out">
                Back
            </span>
        </Link>
    );
};

export default BackToAccount;
