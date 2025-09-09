import { Component, computed, input } from '@angular/core';
import { User } from '../../interfaces/user';

@Component({
    selector: 'user-game',
    templateUrl: 'user-game.component.html',
    styleUrl: 'user-game.component.scss',
    imports: [],
})
export class UserGame {
    userInput = input<User>();

    user = computed(() => this.userInput());
}
