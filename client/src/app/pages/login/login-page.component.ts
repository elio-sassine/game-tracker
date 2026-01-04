import { Component } from '@angular/core';
import { LoginComponent } from '../../components/login/login.component';

@Component({
    selector: 'app-login-page',
    templateUrl: 'login-page.component.html',
    styleUrl: 'login-page.component.scss',
    imports: [LoginComponent],
})
export class LoginPageComponent {}
