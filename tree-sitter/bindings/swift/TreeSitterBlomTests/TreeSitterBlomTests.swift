import XCTest
import SwiftTreeSitter
import TreeSitterBlom

final class TreeSitterBlomTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_blom())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Blom grammar")
    }
}
