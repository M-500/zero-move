import api from '@/utils/ajax'


export const imageCaptchaAPI = (params = {}) => {
	return api.$get("/api/code", params)
}

export const registerAPI = (params = {}) => {
	return api.$post("/sign-up", params)
}


export const pwdLoginAPI = (params = {}) => {
	return api.$post("/api/user/login", params)
}
