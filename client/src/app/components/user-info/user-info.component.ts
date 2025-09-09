import { Component, computed, input } from '@angular/core';
import { User } from '../../interfaces/user';

@Component({
    selector: 'user-info',
    styleUrl: 'user-info.component.scss',
    templateUrl: 'user-info.component.html',
    imports: [],
})
export class UserInfo {
    userInput = input<User>();

    user = computed(() => this.userInput());
}
