//
//  AppUtils.swift
//  GomobilePresentation
//
//  Created by Alexander Timonenkov on 24.05.2023.
//

import Foundation

func mainAsync(_ code: @escaping () -> Void) {
    DispatchQueue.main.async {
        code()
    }
}
