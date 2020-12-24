export const ACT_USER_SET_USER = '[User] Set_User'

const initialState = {
    user: null,
    isLoggedIn: false
};

const authReducer = (state = initialState, action) => {
    switch (action.type) {
        case ACT_USER_SET_USER:
            return {
                user: action.payload,
                isLoggedIn: true
            };
        default:
            return state;
    }
}

export default authReducer;
