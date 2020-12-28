import Head from "next/head";
import {useRouter} from "next/router";
import {useDispatch} from "react-redux";
import {ACT_USER_SET_USER} from "@/domains/auth/reducers/auth";
import {useFormik} from "formik";
import axiosClient from "@/configs/axiosClient";
import * as yup from 'yup';
import ValidationText from "@/domains/share/components/validation-text";
import {isJSON} from "../../helpers/json";

const Auth = () => {
    const router = useRouter()
    const dispatch = useDispatch()
    const formik = useFormik({
        initialValues: {
            email: '',
            password: '',
            passwordConfirm: ''
        },
        onSubmit: data => {
            axiosClient.post('auth/register', data)
                .then(res => {
                    console.log(res, '+++++')
                })
                .catch(err => {
                    const response = err.response.data;

                    const jsonError = isJSON(response.error)
                    jsonError && formik.setErrors(jsonError)
                })
        },
        validationSchema: yup.object().shape({
            email: yup.string().required().email(),
            password: yup.string().required(),
            passwordConfirm: yup.string().required().oneOf([yup.ref('password')])
        })
    })

    return (
        <div className="min-h-screen flex flex-col items-center justify-center bg-gray-300">
            <Head>
                <title>Login</title>
            </Head>

            <div
                className="flex flex-col bg-white shadow-md px-4 sm:px-6 md:px-8 lg:px-10 py-8 rounded-md w-full max-w-md">
                <div className="font-medium self-center text-xl sm:text-2xl uppercase text-gray-800">
                    Register
                </div>
                <div className="mt-10">
                    <form onSubmit={formik.handleSubmit}>
                        <div className="flex flex-col mb-6">
                            <label htmlFor="email" className="mb-1 text-xs sm:text-sm tracking-wide text-gray-600">E-Mail
                                Address:</label>
                            <div className="relative">
                                <div
                                    className="inline-flex items-center justify-center absolute left-0 top-0 h-full w-10 text-gray-400">
                                    <svg className="h-6 w-6" fill="none" strokeLinecap="round" strokeLinejoin="round"
                                         strokeWidth="{2}" viewBox="0 0 24 24" stroke="currentColor">
                                        <path
                                            d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207"/>
                                    </svg>
                                </div>
                                <input
                                    id="email"
                                    name="email"
                                    className="text-sm sm:text-base placeholder-gray-500 pl-10 pr-4 rounded-lg border border-gray-400 w-full py-2 focus:outline-none focus:border-blue-400"
                                    placeholder="E-Mail Address"
                                    onChange={formik.handleChange}
                                    value={formik.values.email}
                                />
                            </div>
                            <ValidationText>{formik.errors.email}</ValidationText>
                        </div>
                        <div className="flex flex-col mb-6">
                            <label htmlFor="password"
                                   className="mb-1 text-xs sm:text-sm tracking-wide text-gray-600">Password:</label>
                            <div className="relative">
                                <div
                                    className="inline-flex items-center justify-center absolute left-0 top-0 h-full w-10 text-gray-400">
                                    <span>
                                        <svg className="h-6 w-6" fill="none" strokeLinecap="round" strokeLinejoin="round"
                                             strokeWidth="{2}" viewBox="0 0 24 24" stroke="currentColor">
                                            <path
                                                d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
                                        </svg>
                                    </span>
                                </div>
                                <input
                                    id="password"
                                    type="password"
                                    name="password"
                                    className="text-sm sm:text-base placeholder-gray-500 pl-10 pr-4 rounded-lg border border-gray-400 w-full py-2 focus:outline-none focus:border-blue-400"
                                    placeholder="Password"
                                    onChange={formik.handleChange}
                                    value={formik.values.password}
                                />
                            </div>
                            <ValidationText>{formik.errors.password}</ValidationText>
                            <div className="relative mt-3">
                                <div
                                    className="inline-flex items-center justify-center absolute left-0 top-0 h-full w-10 text-gray-400">
                                    <span>
                                        <svg className="h-6 w-6" fill="none" strokeLinecap="round" strokeLinejoin="round"
                                             strokeWidth="{2}" viewBox="0 0 24 24" stroke="currentColor">
                                            <path
                                                d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
                                        </svg>
                                    </span>
                                </div>
                                <input
                                    id="passwordConfirm"
                                    type="password"
                                    name="passwordConfirm"
                                    className="text-sm sm:text-base placeholder-gray-500 pl-10 pr-4 rounded-lg border border-gray-400 w-full py-2 focus:outline-none focus:border-blue-400"
                                    placeholder="Password Confirm"
                                    onChange={formik.handleChange}
                                    value={formik.values.passwordConfirm}
                                />
                            </div>
                            <ValidationText>{formik.errors.passwordConfirm}</ValidationText>
                        </div>
                        <div className="flex w-full">
                            <button type="submit"
                                    className="flex items-center justify-center focus:outline-none text-white text-sm sm:text-base bg-blue-600 hover:bg-blue-700 rounded py-2 w-full transition duration-150 ease-in">
                                <span className="mr-2 uppercase">Register</span>
                                <span>
                            <svg className="h-6 w-6" fill="none" strokeLinecap="round" strokeLinejoin="round"
                                 strokeWidth="{2}" viewBox="0 0 24 24" stroke="currentColor">
                                <path d="M13 9l3 3m0 0l-3 3m3-3H8m13 0a9 9 0 11-18 0 9 9 0 0118 0z"/>
                            </svg>
                        </span>
                            </button>
                        </div>
                    </form>
                </div>
                <div className="flex justify-center items-center mt-6">
                    <span
                        onClick={() => router.push('/auth/login')}
                        className="inline-flex items-center font-bold text-blue-500 hover:text-blue-700 text-xs text-center cursor-pointer">
                        <span>
                            <svg className="h-6 w-6" fill="none" strokeLinecap="round" strokeLinejoin="round"
                                 strokeWidth="{2}"
                                 viewBox="0 0 24 24" stroke="currentColor">
                                <path
                                    d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z"/>
                            </svg>
                        </span>
                        <span className="ml-2">Back to login</span>
                    </span>
                </div>
            </div>
        </div>

    )
};

export default Auth