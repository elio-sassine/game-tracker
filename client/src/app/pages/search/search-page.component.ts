import { Component } from '@angular/core';
import { GameHandler } from '../../services/game-fetch.service';
import { SearchComponent } from '../../components/search/search.component';
import { GameComponent } from '../../components/game/game.component';

@Component({
    selector: 'app-search-page',
    templateUrl: './search-page.component.html',
    styleUrl: './search-page.component.scss',
    imports: [SearchComponent, GameComponent],
})
export class SearchPageComponent {
    constructor(public gameHandler: GameHandler) {}
}
