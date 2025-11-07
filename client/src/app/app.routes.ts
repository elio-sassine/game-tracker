import { Routes } from '@angular/router';
import { AppComponent } from './app.component';
import { SearchPage } from './pages/search/search-page.component';
import { RegisterPage } from './pages/register/register-page.component';
import { LoginPage } from './pages/login/login-page.component';
import { UserPage } from './pages/user/user-page.component';

export const routes: Routes = [
    {
        path: '',
        component: AppComponent,
    },
    {
        path: 'search',
        component: SearchPage,
    },
    {
        path: 'register',
        component: RegisterPage,
    },
    {
        path: 'login',
        component: LoginPage,
    },
    {
        path: 'user/:id',
        component: UserPage,
    },
];
