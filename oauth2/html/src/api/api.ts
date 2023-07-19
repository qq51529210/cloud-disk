
const networkError = '网络错误'

export interface error {
    phrase: string;
    detail: string;
}

export const baseURL = (path: string): string => {
    if (process.env.NODE_ENV === 'production') {
        return path
    }
    return '/proxy' + path
}

export const onError = (err: any): error => {
    if (err.response && err.response.data) {
        return <error>err.response.data
    }
    return { phrase: networkError, detail: err.message }
}
