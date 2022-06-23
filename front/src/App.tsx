import React, {useState} from 'react';
import './App.css';
import api from './api'

import {
    Box, Button,
    Container,
    Grid, TextField,
    Typography,
} from '@mui/material';

function App() {
    const [login, setLogin] = useState('');
    const [password, setPassword] = useState('');
    const loginRe = /^[a-zA-Z][a-zA-Z0-9]*$/;
    const loginValid = loginRe.test(login);
    const allFilled = login !== "" && password !== "" && loginValid

    const handleSubmit = (e: React.FormEvent<HTMLButtonElement>) => {
        api.post('/auth/register', {
            login: login,
            password: password,
        })
            .then((res) => {
                console.log(res)
            }).catch(err => console.log(err))
    }

    return (
        <Container maxWidth={'xs'}>
            <Box sx={{m: '20%'}}>
                <Typography
                    component={'h1'}
                    variant={'h5'}
                    align={'center'}
                    marginBottom={5}
                >
                    Регистрация
                </Typography>
                <Grid container spacing={2} mt={5} mb={5}>
                    <Grid item xs={12} sm={12}>
                        <TextField
                            fullWidth={true}
                            required={true}
                            label={'Login'}
                            value={login}
                            error={!loginValid && login !== ""}
                            onChange={(e) => {
                                e.preventDefault()
                                setLogin(e.target.value)
                            }}
                        />
                    </Grid>
                    <Grid item xs={12} sm={12}>
                        <TextField
                            fullWidth={true}
                            required={true}
                            label={'Password'}
                            value={password}
                            onChange={(e) => {
                                e.preventDefault()
                                setPassword(e.target.value)
                            }}
                        />
                    </Grid>
                </Grid>
                <Button
                    onClick={handleSubmit}
                    disabled={!allFilled}
                    fullWidth={true}
                    sx={{mt: '5'}}
                >
                    Создать аккаунт
                </Button>
            </Box>
        </Container>
    );
}

export default App;
