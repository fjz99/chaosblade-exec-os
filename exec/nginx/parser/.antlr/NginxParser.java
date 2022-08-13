// Generated from /chaos/chaosblade-exec-os/exec/nginx/parser/Nginx.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class NginxParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, Value=15, STR_EXT=16, 
		LINE_COMMENT=17, REGEXP_PREFIXED=18, QUOTED_STRING=19, SINGLE_QUOTED=20, 
		WS=21;
	public static final int
		RULE_config = 0, RULE_statement = 1, RULE_genericStatement = 2, RULE_regexHeaderStatement = 3, 
		RULE_block = 4, RULE_genericBlockHeader = 5, RULE_ifStatement = 6, RULE_ifBody = 7, 
		RULE_regexp = 8, RULE_locationBlockHeader = 9, RULE_rewriteStatement = 10;
	private static String[] makeRuleNames() {
		return new String[] {
			"config", "statement", "genericStatement", "regexHeaderStatement", "block", 
			"genericBlockHeader", "ifStatement", "ifBody", "regexp", "locationBlockHeader", 
			"rewriteStatement"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "';'", "'{'", "'}'", "'if'", "'('", "')'", "'\\.'", "'^'", "'location'", 
			"'rewrite'", "'last'", "'break'", "'redirect'", "'permanent'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, "Value", "STR_EXT", "LINE_COMMENT", "REGEXP_PREFIXED", 
			"QUOTED_STRING", "SINGLE_QUOTED", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Nginx.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public NginxParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ConfigContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public ConfigContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_config; }
	}

	public final ConfigContext config() throws RecognitionException {
		ConfigContext _localctx = new ConfigContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_config);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(24); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				setState(24);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
				case 1:
					{
					setState(22);
					statement();
					}
					break;
				case 2:
					{
					setState(23);
					block();
					}
					break;
				}
				}
				setState(26); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << Value) | (1L << REGEXP_PREFIXED))) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public RewriteStatementContext rewriteStatement() {
			return getRuleContext(RewriteStatementContext.class,0);
		}
		public GenericStatementContext genericStatement() {
			return getRuleContext(GenericStatementContext.class,0);
		}
		public RegexHeaderStatementContext regexHeaderStatement() {
			return getRuleContext(RegexHeaderStatementContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(31);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__9:
				{
				setState(28);
				rewriteStatement();
				}
				break;
			case Value:
				{
				setState(29);
				genericStatement();
				}
				break;
			case REGEXP_PREFIXED:
				{
				setState(30);
				regexHeaderStatement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(33);
			match(T__0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class GenericStatementContext extends ParserRuleContext {
		public List<TerminalNode> Value() { return getTokens(NginxParser.Value); }
		public TerminalNode Value(int i) {
			return getToken(NginxParser.Value, i);
		}
		public List<RegexpContext> regexp() {
			return getRuleContexts(RegexpContext.class);
		}
		public RegexpContext regexp(int i) {
			return getRuleContext(RegexpContext.class,i);
		}
		public GenericStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_genericStatement; }
	}

	public final GenericStatementContext genericStatement() throws RecognitionException {
		GenericStatementContext _localctx = new GenericStatementContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_genericStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(35);
			match(Value);
			setState(40);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__4) | (1L << T__6) | (1L << T__7) | (1L << Value))) != 0)) {
				{
				setState(38);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
				case 1:
					{
					setState(36);
					match(Value);
					}
					break;
				case 2:
					{
					setState(37);
					regexp();
					}
					break;
				}
				}
				setState(42);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RegexHeaderStatementContext extends ParserRuleContext {
		public TerminalNode REGEXP_PREFIXED() { return getToken(NginxParser.REGEXP_PREFIXED, 0); }
		public TerminalNode Value() { return getToken(NginxParser.Value, 0); }
		public RegexHeaderStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_regexHeaderStatement; }
	}

	public final RegexHeaderStatementContext regexHeaderStatement() throws RecognitionException {
		RegexHeaderStatementContext _localctx = new RegexHeaderStatementContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_regexHeaderStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(43);
			match(REGEXP_PREFIXED);
			setState(44);
			match(Value);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BlockContext extends ParserRuleContext {
		public LocationBlockHeaderContext locationBlockHeader() {
			return getRuleContext(LocationBlockHeaderContext.class,0);
		}
		public GenericBlockHeaderContext genericBlockHeader() {
			return getRuleContext(GenericBlockHeaderContext.class,0);
		}
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public List<IfStatementContext> ifStatement() {
			return getRuleContexts(IfStatementContext.class);
		}
		public IfStatementContext ifStatement(int i) {
			return getRuleContext(IfStatementContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(48);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__8:
				{
				setState(46);
				locationBlockHeader();
				}
				break;
			case Value:
				{
				setState(47);
				genericBlockHeader();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(50);
			match(T__1);
			setState(56);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__3) | (1L << T__8) | (1L << T__9) | (1L << Value) | (1L << REGEXP_PREFIXED))) != 0)) {
				{
				setState(54);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
				case 1:
					{
					setState(51);
					statement();
					}
					break;
				case 2:
					{
					setState(52);
					block();
					}
					break;
				case 3:
					{
					setState(53);
					ifStatement();
					}
					break;
				}
				}
				setState(58);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(59);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class GenericBlockHeaderContext extends ParserRuleContext {
		public List<TerminalNode> Value() { return getTokens(NginxParser.Value); }
		public TerminalNode Value(int i) {
			return getToken(NginxParser.Value, i);
		}
		public List<RegexpContext> regexp() {
			return getRuleContexts(RegexpContext.class);
		}
		public RegexpContext regexp(int i) {
			return getRuleContext(RegexpContext.class,i);
		}
		public GenericBlockHeaderContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_genericBlockHeader; }
	}

	public final GenericBlockHeaderContext genericBlockHeader() throws RecognitionException {
		GenericBlockHeaderContext _localctx = new GenericBlockHeaderContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_genericBlockHeader);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(61);
			match(Value);
			setState(66);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__4) | (1L << T__6) | (1L << T__7) | (1L << Value))) != 0)) {
				{
				setState(64);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
				case 1:
					{
					setState(62);
					match(Value);
					}
					break;
				case 2:
					{
					setState(63);
					regexp();
					}
					break;
				}
				}
				setState(68);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IfStatementContext extends ParserRuleContext {
		public IfBodyContext ifBody() {
			return getRuleContext(IfBodyContext.class,0);
		}
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_ifStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(69);
			match(T__3);
			setState(70);
			ifBody();
			setState(71);
			match(T__1);
			setState(75);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__9) | (1L << Value) | (1L << REGEXP_PREFIXED))) != 0)) {
				{
				{
				setState(72);
				statement();
				}
				}
				setState(77);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(78);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IfBodyContext extends ParserRuleContext {
		public List<TerminalNode> Value() { return getTokens(NginxParser.Value); }
		public TerminalNode Value(int i) {
			return getToken(NginxParser.Value, i);
		}
		public RegexpContext regexp() {
			return getRuleContext(RegexpContext.class,0);
		}
		public IfBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifBody; }
	}

	public final IfBodyContext ifBody() throws RecognitionException {
		IfBodyContext _localctx = new IfBodyContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_ifBody);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(80);
			match(T__4);
			setState(81);
			match(Value);
			setState(83);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				{
				setState(82);
				match(Value);
				}
				break;
			}
			setState(87);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				{
				setState(85);
				match(Value);
				}
				break;
			case 2:
				{
				setState(86);
				regexp();
				}
				break;
			}
			setState(89);
			match(T__5);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RegexpContext extends ParserRuleContext {
		public List<TerminalNode> Value() { return getTokens(NginxParser.Value); }
		public TerminalNode Value(int i) {
			return getToken(NginxParser.Value, i);
		}
		public List<RegexpContext> regexp() {
			return getRuleContexts(RegexpContext.class);
		}
		public RegexpContext regexp(int i) {
			return getRuleContext(RegexpContext.class,i);
		}
		public RegexpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_regexp; }
	}

	public final RegexpContext regexp() throws RecognitionException {
		RegexpContext _localctx = new RegexpContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_regexp);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(98); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					setState(98);
					_errHandler.sync(this);
					switch (_input.LA(1)) {
					case T__6:
						{
						setState(91);
						match(T__6);
						}
						break;
					case T__7:
						{
						setState(92);
						match(T__7);
						}
						break;
					case Value:
						{
						setState(93);
						match(Value);
						}
						break;
					case T__4:
						{
						setState(94);
						match(T__4);
						setState(95);
						regexp();
						setState(96);
						match(T__5);
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(100); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LocationBlockHeaderContext extends ParserRuleContext {
		public List<TerminalNode> Value() { return getTokens(NginxParser.Value); }
		public TerminalNode Value(int i) {
			return getToken(NginxParser.Value, i);
		}
		public RegexpContext regexp() {
			return getRuleContext(RegexpContext.class,0);
		}
		public LocationBlockHeaderContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_locationBlockHeader; }
	}

	public final LocationBlockHeaderContext locationBlockHeader() throws RecognitionException {
		LocationBlockHeaderContext _localctx = new LocationBlockHeaderContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_locationBlockHeader);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(102);
			match(T__8);
			setState(104);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				{
				setState(103);
				match(Value);
				}
				break;
			}
			setState(108);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				{
				setState(106);
				match(Value);
				}
				break;
			case 2:
				{
				setState(107);
				regexp();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RewriteStatementContext extends ParserRuleContext {
		public List<TerminalNode> Value() { return getTokens(NginxParser.Value); }
		public TerminalNode Value(int i) {
			return getToken(NginxParser.Value, i);
		}
		public RegexpContext regexp() {
			return getRuleContext(RegexpContext.class,0);
		}
		public RewriteStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rewriteStatement; }
	}

	public final RewriteStatementContext rewriteStatement() throws RecognitionException {
		RewriteStatementContext _localctx = new RewriteStatementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_rewriteStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(110);
			match(T__9);
			setState(113);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				{
				setState(111);
				match(Value);
				}
				break;
			case 2:
				{
				setState(112);
				regexp();
				}
				break;
			}
			setState(115);
			match(Value);
			setState(117);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13))) != 0)) {
				{
				setState(116);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13))) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\27z\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t\13\4"+
		"\f\t\f\3\2\3\2\6\2\33\n\2\r\2\16\2\34\3\3\3\3\3\3\5\3\"\n\3\3\3\3\3\3"+
		"\4\3\4\3\4\7\4)\n\4\f\4\16\4,\13\4\3\5\3\5\3\5\3\6\3\6\5\6\63\n\6\3\6"+
		"\3\6\3\6\3\6\7\69\n\6\f\6\16\6<\13\6\3\6\3\6\3\7\3\7\3\7\7\7C\n\7\f\7"+
		"\16\7F\13\7\3\b\3\b\3\b\3\b\7\bL\n\b\f\b\16\bO\13\b\3\b\3\b\3\t\3\t\3"+
		"\t\5\tV\n\t\3\t\3\t\5\tZ\n\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\6\ne"+
		"\n\n\r\n\16\nf\3\13\3\13\5\13k\n\13\3\13\3\13\5\13o\n\13\3\f\3\f\3\f\5"+
		"\ft\n\f\3\f\3\f\5\fx\n\f\3\f\2\2\r\2\4\6\b\n\f\16\20\22\24\26\2\3\3\2"+
		"\r\20\2\u0086\2\32\3\2\2\2\4!\3\2\2\2\6%\3\2\2\2\b-\3\2\2\2\n\62\3\2\2"+
		"\2\f?\3\2\2\2\16G\3\2\2\2\20R\3\2\2\2\22d\3\2\2\2\24h\3\2\2\2\26p\3\2"+
		"\2\2\30\33\5\4\3\2\31\33\5\n\6\2\32\30\3\2\2\2\32\31\3\2\2\2\33\34\3\2"+
		"\2\2\34\32\3\2\2\2\34\35\3\2\2\2\35\3\3\2\2\2\36\"\5\26\f\2\37\"\5\6\4"+
		"\2 \"\5\b\5\2!\36\3\2\2\2!\37\3\2\2\2! \3\2\2\2\"#\3\2\2\2#$\7\3\2\2$"+
		"\5\3\2\2\2%*\7\21\2\2&)\7\21\2\2\')\5\22\n\2(&\3\2\2\2(\'\3\2\2\2),\3"+
		"\2\2\2*(\3\2\2\2*+\3\2\2\2+\7\3\2\2\2,*\3\2\2\2-.\7\24\2\2./\7\21\2\2"+
		"/\t\3\2\2\2\60\63\5\24\13\2\61\63\5\f\7\2\62\60\3\2\2\2\62\61\3\2\2\2"+
		"\63\64\3\2\2\2\64:\7\4\2\2\659\5\4\3\2\669\5\n\6\2\679\5\16\b\28\65\3"+
		"\2\2\28\66\3\2\2\28\67\3\2\2\29<\3\2\2\2:8\3\2\2\2:;\3\2\2\2;=\3\2\2\2"+
		"<:\3\2\2\2=>\7\5\2\2>\13\3\2\2\2?D\7\21\2\2@C\7\21\2\2AC\5\22\n\2B@\3"+
		"\2\2\2BA\3\2\2\2CF\3\2\2\2DB\3\2\2\2DE\3\2\2\2E\r\3\2\2\2FD\3\2\2\2GH"+
		"\7\6\2\2HI\5\20\t\2IM\7\4\2\2JL\5\4\3\2KJ\3\2\2\2LO\3\2\2\2MK\3\2\2\2"+
		"MN\3\2\2\2NP\3\2\2\2OM\3\2\2\2PQ\7\5\2\2Q\17\3\2\2\2RS\7\7\2\2SU\7\21"+
		"\2\2TV\7\21\2\2UT\3\2\2\2UV\3\2\2\2VY\3\2\2\2WZ\7\21\2\2XZ\5\22\n\2YW"+
		"\3\2\2\2YX\3\2\2\2YZ\3\2\2\2Z[\3\2\2\2[\\\7\b\2\2\\\21\3\2\2\2]e\7\t\2"+
		"\2^e\7\n\2\2_e\7\21\2\2`a\7\7\2\2ab\5\22\n\2bc\7\b\2\2ce\3\2\2\2d]\3\2"+
		"\2\2d^\3\2\2\2d_\3\2\2\2d`\3\2\2\2ef\3\2\2\2fd\3\2\2\2fg\3\2\2\2g\23\3"+
		"\2\2\2hj\7\13\2\2ik\7\21\2\2ji\3\2\2\2jk\3\2\2\2kn\3\2\2\2lo\7\21\2\2"+
		"mo\5\22\n\2nl\3\2\2\2nm\3\2\2\2o\25\3\2\2\2ps\7\f\2\2qt\7\21\2\2rt\5\22"+
		"\n\2sq\3\2\2\2sr\3\2\2\2tu\3\2\2\2uw\7\21\2\2vx\t\2\2\2wv\3\2\2\2wx\3"+
		"\2\2\2x\27\3\2\2\2\25\32\34!(*\628:BDMUYdfjnsw";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}