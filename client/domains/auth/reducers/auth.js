const authReducer = (state = {email: 'test'}, action) => {
    switch (action.type) {
        case 'SET_AUTH_USER':
            return {
                email: 'test',
                name: 'test'
            };
        default:
            return state;
    }
}

export default authReducer;