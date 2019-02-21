package freerangeserver

type GameContext struct {
	levelmanager  *LevelManager
	level         *Level
	levelViewPort *LevelViewPort
	levelFactory  LevelFactory
	entityFactory EntityFactory
}

func NewGameContext(levelmanager *LevelManager, levelFactory LevelFactory, entityFactory EntityFactory) *GameContext {
	c := new(GameContext)
	c.levelmanager = levelmanager
	c.levelFactory = levelFactory
	c.entityFactory = entityFactory
	return c
}
func (gamecontext *GameContext) LoadLevel(levelID int64) {
	if gamecontext.level != nil {
		gamecontext.levelmanager.CloseLevel(gamecontext.level)
	}
	gamecontext.level = gamecontext.levelmanager.GetLevel(levelID, gamecontext.levelFactory, gamecontext.entityFactory)
	gamecontext.levelViewPort = NewLevelViewPort()

}
func (gamecontext *GameContext) Refresh() (created []Entity, destroyed []int64, moved []Position) {
	result := gamecontext.levelViewPort.Refresh(gamecontext.level)
	return result.created, result.destroyed, result.moved
}

func (gamecontext *GameContext) ClickAction(entityID int64) {
	e := gamecontext.level.GetEntity(entityID)
	e.clickAction(gamecontext.level, gamecontext.levelViewPort)
}

//LoadAssets loads the assets needed to render the game state
func (gamecontext *GameContext) LoadAssets() []byte {
	return []byte(`
		{ 
			"images": {
				"bg": "https://twemoji.maxcdn.com/72x72/1f306.png",
				"player": "https://twemoji.maxcdn.com/2/72x72/1f600.png",
				"ground": "assets/platform.png",
				"house": "https://twemoji.maxcdn.com/2/72x72/1f3d8.png",
				"hospital": "https://twemoji.maxcdn.com/2/72x72/1f3e5.png",
				"npc": "assets/face-positive/beaming face with smiling eyes.png"
			}
		}`)
}
