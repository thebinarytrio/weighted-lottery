const initialState = {
	isMenuOn: false,
	
}

export default (state = initialState, action) => {
	switch (action.type) {
		case "isMenuOn":
			return Object.assign({}, state, { isMenuOn: true })
		case "isMenuOff":
			return Object.assign({}, state, { isMenuOn: false })
		default:
			return state
	}
}