//
//  FileManager+Extensions.swift
//  GomobilePresentation
//
//  Created by Alexander Timonenkov on 23.05.2023.
//

import Foundation

extension FileManager {
    static var engineAssetsDirectory: URL? {
        Bundle.main.url(forResource: "EngineAssets", withExtension: nil)
    }

    static var engineAssetsDirectoryForEngine: String {
        Self.engineAssetsDirectory?
            .absoluteString
            .replacingOccurrences(of: "file://", with: "") ?? ""
    }
}
