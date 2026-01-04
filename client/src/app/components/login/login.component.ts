import { Component } from '@angular/core';
import {
    FormControl,
    FormGroup,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { MatFormField } from '@angular/material/form-field';
import { LoginHandler } from '../../services/login-handler.service';

@Component({
    selector: 'app-login-component',
    templateUrl: 'login.component.html',
    styleUrl: 'login.component.scss',
    imports: [
        MatFormField,
        MatInputModule,
        ReactiveFormsModule,
        MatButtonModule,
    ],
})
export class LoginComponent {
    constructor(private loginHandler: LoginHandler) {}

    loginForm = new FormGroup({
        email: new FormControl('', [Validators.required, Validators.email]),
        password: new FormControl('', [
            Validators.required,
            Validators.minLength(6),
        ]),
    });

    get email() {
        return this.loginForm.get('email');
    }

    get password() {
        return this.loginForm.get('password');
    }

    onSubmit() {
        if (this.loginForm.value.email && this.loginForm.value.password) {
            this.loginHandler.loginUser(
                this.loginForm.value.email,
                this.loginForm.value.password
            );
        }
    }
}
