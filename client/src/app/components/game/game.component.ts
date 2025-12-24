import { Component, computed, input } from '@angular/core';
import { GameHandler } from '../../services/game-fetch.service';
import { GameTrackButtonComponent } from '../game-track-button/game-track-button.component';
import { Game } from '../../interfaces/game';
import { MatCardModule } from '@angular/material/card';

@Component({
    selector: 'game',
    templateUrl: './game.component.html',
    styleUrl: './game.component.scss',
    imports: [MatCardModule, GameTrackButtonComponent],
})
export class GameComponent {
    gameInput = input<Game>();

    game = computed(() => this.gameInput());
}
