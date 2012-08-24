class ReplyListViewController extends KDListViewController
  constructor:->
    super
    @_hasBackgrounActivity = no
    @startListeners()

  instantiateListItems:(items, keepDeletedComments = no)->

    newItems = []

    items.sort (a,b) =>
      a = a.meta.createdAt
      b = b.meta.createdAt
      if a<b then -1 else if a>b then 1 else 0

    for comment, i in items
      nextComment = items[i+1]

      skipComment = no
      if nextComment? and comment.deletedAt
        if Date.parse(nextComment.meta.createdAt) > Date.parse(comment.deletedAt)
          skipComment = yes

      if not nextComment and comment.deletedAt
        skipComment = yes

      skipComment = no if keepDeletedComments

      unless skipComment
        commentView = @getListView().addItem comment
        newItems.push commentView

    return newItems

  startListeners:->
    listView = @getListView()

    listView.on 'ItemWasAdded', (view, index)=>
      view.on 'CommentIsDeleted', ->
        listView.emit "CommentIsDeleted"

    listView.on "AllCommentsLinkWasClicked", (commentHeader)=>

      return if @_hasBackgrounActivity

      # some problems when logged out server doesnt responds
      @utils.wait 5000, -> listView.emit "BackgroundActivityFinished"

      {meta} = listView.getData()

      listView.emit "BackgroundActivityStarted"
      @_hasBackgrounActivity = yes
      @_removedBefore = no
      @fetchRelativeComments 10, meta.createdAt

    listView.registerListener
      KDEventTypes  : "CommentSubmitted"
      listener      : @
      callback      : (pubInst, reply)->
        model = listView.getData()
        listView.emit "BackgroundActivityStarted"
        model.reply reply, (err, reply)->
          # listView.emit "AllCommentsLinkWasClicked"
          listView.addItem reply
          listView.emit "OwnCommentHasArrived"
          listView.emit "BackgroundActivityFinished"

  fetchCommentsByRange:(from,to,callback)=>
    [to,callback] = [callback,to] unless callback
    query = {from,to}
    message = @getListView().getData()

    log "reply fcbr", message

    message.commentsByRange query,(err,comments)=>
      @getListView().emit "BackgroundActivityFinished"
      callback err,comments

  fetchAllComments:(skipCount=3, callback = noop)=>

    log "reply fac"

    listView = @getListView()
    listView.emit "BackgroundActivityStarted"
    message = @getListView().getData()
    message.restComments skipCount, (err, comments)=>

      log "fetched all comments with err ", err, comments

      listView.emit "BackgroundActivityFinished"
      listView.emit "AllCommentsWereAdded"
      callback err, comments

  fetchRelativeComments:(_limit = 10, _after)=>
    listView = @getListView()
    message = @getListView().getData()

    log "reply frc", message


    message.fetchRelativeComments limit:_limit, after:_after, (err, comments)=>

      log "fetched relative comments with err", err, comments

      if not @_removedBefore
        @removeAllItems()
        @_removedBefore = yes
      log "before reply inst"
      @instantiateListItems comments[_limit-10...], yes
      log "after reply inst"

      if comments.length is _limit
        startTime = comments[comments.length-1].meta.createdAt
        @fetchRelativeComments 11, startTime
      else
        listView = @getListView()
        listView.emit "BackgroundActivityFinished"
        listView.emit "AllCommentsWereAdded"
        @_hasBackgrounActivity = no

  replaceAllComments:(comments)->
    log "replaceAllComments"
    @removeAllItems()
    @instantiateListItems comments
