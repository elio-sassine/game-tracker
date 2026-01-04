import { Component } from '@angular/core';
import {
    FormControl,
    FormGroup,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormField } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { RegistrationHandler } from '../../services/registeration-handler.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
    selector: 'app-register-component',
    templateUrl: 'register.component.html',
    styleUrl: 'register.component.scss',
    imports: [
        MatFormField,
        MatInputModule,
        ReactiveFormsModule,
        MatButtonModule,
    ],
})
export class RegisterComponent {
    constructor(
        private registrationHandler: RegistrationHandler,
        private snackBar: MatSnackBar
    ) {}

    registerForm = new FormGroup({
        email: new FormControl('', [Validators.email, Validators.required]),
        username: new FormControl('', Validators.required),
        password: new FormControl('', [
            Validators.required,
            Validators.minLength(6),
        ]),
    });

    get email() {
        return this.registerForm.get('email');
    }

    get password() {
        return this.registerForm.get('password');
    }

    onSubmit() {
        if (
            this.registerForm.value.email &&
            this.registerForm.value.password &&
            this.registerForm.value.username
        ) {
            this.registrationHandler.makeUser(
                this.registerForm.value.email,
                this.registerForm.value.password,
                this.registerForm.value.username
            );
        }
    }
}
