import { Component, computed, input } from '@angular/core';
import { User } from '../../interfaces/user';

@Component({
    selector: 'app-user-info',
    styleUrl: 'user-info.component.scss',
    templateUrl: 'user-info.component.html',
    imports: [],
})
export class UserInfoComponent {
    userInput = input<User>();

    user = computed(() => this.userInput());
}
