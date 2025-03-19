import { createContext, useContext, useState } from 'react';

const AuthContext = createContext();

export function AuthProvider({ children }) {
	const [user, setUser] = useState(null);
	const [activeSession, setActiveSession] = useState(null);

	return (
		<AuthContext.Provider value={{ user, setUser, activeSession, setActiveSession }}>
			{children}
		</AuthContext.Provider>
	);
}

export function useAuth() {
	return useContext(AuthContext);
}
