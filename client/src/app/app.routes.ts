import { Routes } from '@angular/router';
import { AppComponent } from './app.component';
import { SearchPageComponent } from './pages/search/search-page.component';
import { RegisterPageComponent } from './pages/register/register-page.component';
import { LoginPageComponent } from './pages/login/login-page.component';
import { UserPageComponent } from './pages/user/user-page.component';

export const routes: Routes = [
    {
        path: '',
        component: AppComponent,
    },
    {
        path: 'search',
        component: SearchPageComponent,
    },
    {
        path: 'register',
        component: RegisterPageComponent,
    },
    {
        path: 'login',
        component: LoginPageComponent,
    },
    {
        path: 'user/:id',
        component: UserPageComponent,
    },
];
