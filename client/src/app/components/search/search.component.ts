import { Component } from '@angular/core';
import {
    FormControl,
    FormGroup,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { MatFormField } from '@angular/material/form-field';
import { GameHandler } from '../../services/game-fetch.service';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';

@Component({
    selector: 'app-search',
    templateUrl: './search.component.html',
    styleUrl: './search.component.scss',
    imports: [
        ReactiveFormsModule,
        MatFormField,
        MatInputModule,
        MatButtonModule,
    ],
})
export class SearchComponent {
    constructor(private gameHandler: GameHandler) {}

    nameForm = new FormGroup({
        name: new FormControl('', Validators.required),
    });

    onSubmit() {
        this.gameHandler.getGames(this.nameForm.value.name);
        this.nameForm.reset();
    }
}
